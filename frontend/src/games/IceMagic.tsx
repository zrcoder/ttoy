import { Events } from "@wailsio/runtime";
import { Button, message, Select } from "antd";
import React, { useEffect, useRef, useState } from "react";
import {
  Chapter,
  Sprite,
} from "../../bindings/github.com/zrcoder/ttoy/game/icemagic";
import * as Game from "../../bindings/github.com/zrcoder/ttoy/game/icemagic/game";
import { contentHeight } from "../components/common/layout";

type State = "succeed" | "failed" | "playing";

const cellStyle = (cell: Sprite["Cell"]) => ({
  width: 32,
  height: 32,
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
  backgroundColor: "transparent",
  borderTop: "2px solid" + (cell?.BorderTop ? " white" : " transparent"),
  borderBottom: "2px solid" + (cell?.BorderBottom ? " white" : " transparent"),
  borderLeft: "2px solid" + (cell?.BorderLeft ? " white" : " transparent"),
  borderRight: "2px solid" + (cell?.BorderRight ? " white" : " transparent"),
  userSelect: "none" as const,
});

const btnCol = {
  display: "flex",
  flexDirection: "column" as const,
  alignItems: "center",
  gap: 4,
};

const IceMagic: React.FC = () => {
  const [grid, setGrid] = useState<Sprite[][] | null>(null);
  const [state, setState] = useState<State>("playing");
  const [chapters, setChapters] = useState<Chapter[]>([]);
  const [selectedChapter, setSelectedChapter] = useState(1);
  const [selectedLevel, setSelectedLevel] = useState(1);
  const [messageApi, contextHolder] = message.useMessage();

  const chapterRef = useRef(selectedChapter);

  useEffect(() => {
    Game.Chapters().then((c) => {
      const cs = (c || []) as Chapter[];
      setChapters(cs);
      Game.SelectLevel(1, 1).then(() =>
        Game.Grid().then((g) => setGrid(g as Sprite[][] | null)),
      );
    });

    const stop = Events.On(
      "icemagic:update",
      (ev: { data: { Grid: Sprite[][]; State: State } }) => {
        setGrid(ev.data.Grid);
        setState(ev.data.State);
      },
    );
    return stop;
  }, []);

  useEffect(() => {
    if (state === "succeed") messageApi.success("Congratulations! You won!");
    else if (state === "failed")
      messageApi.error("Game Over! You were burned!");
  }, [state, messageApi]);

  useEffect(() => {
    const onKey = (e: KeyboardEvent) => {
      const actions: Record<string, () => void> = {
        J: Game.MoveLeft,
        L: Game.MoveRight,
        A: Game.MagicLeft,
        D: Game.MagicRight,
      };
      actions[e.key.toUpperCase()]?.();
    };
    window.addEventListener("keydown", onKey);
    return () => window.removeEventListener("keydown", onKey);
  }, []);

  const blur = () => (document.activeElement as HTMLElement)?.blur();

  const selectLevel = (chapter: number, level: number) => {
    Game.SelectLevel(chapter, level).then(() => {
      Game.Grid().then((g) => setGrid(g as Sprite[][] | null));
      setState("playing");
      blur();
    });
  };

  const handleChapterChange = (v: number) => {
    chapterRef.current = v;
    setSelectedChapter(v);
    setSelectedLevel(1);
    selectLevel(v, 1);
  };

  const handleLevelChange = (v: number) => {
    setSelectedLevel(v);
    selectLevel(chapterRef.current, v);
  };

  const handleReset = () => {
    Game.Reset().then(() => {
      Game.Grid().then((g) => setGrid(g as Sprite[][] | null));
      setState("playing");
      blur();
    });
  };

  const handleNext = () => {
    const chapterCount = chapters.length;
    const levelCount = chapters[selectedChapter - 1] || 1;
    let nextC = selectedChapter;
    let nextL = selectedLevel + 1;
    if (nextL > levelCount) {
      nextL = 1;
      nextC = selectedChapter + 1;
      if (nextC > chapterCount) {
        nextC = 1;
        nextL = 1;
      }
    }
    setSelectedChapter(nextC);
    setSelectedLevel(nextL);
    chapterRef.current = nextC;
    selectLevel(nextC, nextL);
  };

  const renderCell = (sprite: Sprite, row: number, col: number) => {
    const cell = sprite.Cell;
    const img = cell?.Images?.[0];
    return (
      <div key={`${row}-${col}`} style={cellStyle(cell)}>
        {img && <img src={img} alt="" style={{ width: 32, height: 32 }} />}
      </div>
    );
  };

  const renderGrid = () => {
    if (!grid || grid.length === 0) {
      return (
        <div
          style={{
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            height: contentHeight,
          }}
        >
          Loading...
        </div>
      );
    }
    return (
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          backgroundColor: "lightblue",
        }}
      >
        {grid.map((row, ri) => (
          <div key={ri} style={{ display: "flex" }}>
            {row?.map((sprite, ci) => renderCell(sprite, ri, ci))}
          </div>
        ))}
      </div>
    );
  };

  const levelOptions = Array.from(
    { length: chapters[selectedChapter - 1] || 1 },
    (_, i) => ({
      value: i + 1,
      label: `${selectedChapter}-${i + 1}`,
    }),
  );

  const levelBar = (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        gap: 8,
        marginBottom: 8,
      }}
    >
      <span
        style={{
          lineHeight: "32px",
          fontWeight: "bold",
          fontSize: 14,
          letterSpacing: 1,
          marginRight: 8,
        }}
      >
        LEVEL
      </span>
      <Select
        value={selectedChapter}
        onChange={handleChapterChange}
        style={{ width: 80 }}
        listHeight={320}
        options={chapters.map((_, i) => ({
          value: i + 1,
          label: `Ch ${i + 1}`,
        }))}
      />
      <Select
        value={selectedLevel}
        onChange={handleLevelChange}
        style={{ width: 80 }}
        listHeight={320}
        options={levelOptions}
      />
      <Button
        style={{ marginLeft: 16 }}
        onClick={() => {
          handleReset();
          blur();
        }}
      >
        Reset
      </Button>
      <Button onClick={handleNext}>Next</Button>
    </div>
  );

  const controlPad = (
    <div style={{ display: "flex", gap: 20 }}>
      {[
        { key: "A", icon: "↙", action: Game.MagicLeft },
        { key: "J", icon: "←", action: Game.MoveLeft },
        { key: "L", icon: "→", action: Game.MoveRight },
        { key: "D", icon: "↘", action: Game.MagicRight },
      ].map(({ key, icon, action }) => (
        <div key={key} style={btnCol}>
          <Button size="large" onClick={action}>
            {icon}
          </Button>
          <span style={{ fontSize: 12, color: "#666" }}>{key}</span>
        </div>
      ))}
    </div>
  );

  return (
    <div
      style={{
        height: "100%",
        display: "flex",
        flexDirection: "column",
        padding: 16,
      }}
    >
      {contextHolder}
      <div
        style={{
          flex: 1,
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          alignItems: "center",
          gap: 32,
        }}
      >
        {levelBar}
        {renderGrid()}
        {controlPad}
      </div>
    </div>
  );
};

export default IceMagic;
