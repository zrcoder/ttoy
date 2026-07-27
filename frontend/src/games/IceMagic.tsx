import { Events } from "@wailsio/runtime";
import { Button, message, Typography } from "antd";
import React, { useEffect, useState } from "react";
import * as Game from "../../bindings/github.com/zrcoder/ttoy/game/icemagic/game";
import { Sprite } from "../../bindings/github.com/zrcoder/ttoy/game/icemagic";

import { contentHeight } from "../components/common/layout";

type State = "succeed" | "failed" | "playing";

const IceMagic: React.FC = () => {
  const [grid, setGrid] = useState<Sprite[][] | null>(null);
  const [state, setState] = useState<State>("playing");
  const [messageApi, contextHolder] = message.useMessage();

  useEffect(() => {
    Game.Grid().then((g) => setGrid(g as Sprite[][] | null));

    const stop = Events.On(
      "icemagic:update",
      (ev: { data: { Grid: Sprite[][]; State: State } }) => {
        setGrid(ev.data.Grid);
        setState(ev.data.State);
      },
    );

    return () => {
      stop();
    };
  }, []);

  useEffect(() => {
    if (state === "succeed") {
      messageApi.success("Congratulations! You won!");
    } else if (state === "failed") {
      messageApi.error("Game Over! You were burned!");
    }
  }, [state, messageApi]);

  const handleKeyDown = (e: KeyboardEvent) => {
    const key = e.key.toUpperCase();
    switch (key) {
      case "J":
        Game.MoveLeft();
        break;
      case "L":
        Game.MoveRight();
        break;
      case "A":
        Game.MagicLeft();
        break;
      case "D":
        Game.MagicRight();
        break;
    }
  };

  useEffect(() => {
    window.addEventListener("keydown", handleKeyDown);
    return () => {
      window.removeEventListener("keydown", handleKeyDown);
    };
  }, []);

  const renderCell = (sprite: Sprite, rowIndex: number, colIndex: number) => {
    const cell = sprite.Cell;
    const imgPath = cell?.Images?.[0];

    return (
      <div
        key={`${rowIndex}-${colIndex}`}
        style={{
          width: 32,
          height: 32,
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          backgroundColor: "transparent",
          borderTop: cell?.BorderTop
            ? "1px solid rgba(255,255,255,0.5)"
            : "none",
          borderBottom: cell?.BorderBottom
            ? "1px solid rgba(255,255,255,0.5)"
            : "none",
          borderLeft: cell?.BorderLeft
            ? "1px solid rgba(255,255,255,0.5)"
            : "none",
          borderRight: cell?.BorderRight
            ? "1px solid rgba(255,255,255,0.5)"
            : "none",
          userSelect: "none",
        }}
      >
        {imgPath && (
          <img src={imgPath} alt="" style={{ width: 32, height: 32 }} />
        )}
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
          backgroundColor: "white",
        }}
      >
        {grid.map((row, rowIndex) => (
          <div key={rowIndex} style={{ display: "flex" }}>
            {row?.map((sprite, colIndex) =>
              renderCell(sprite, rowIndex, colIndex),
            )}
          </div>
        ))}
      </div>
    );
  };

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
      <Typography.Title
        level={4}
        style={{ textAlign: "center", margin: "0 0 16px" }}
      >
        IceMagic
      </Typography.Title>
      <div
        style={{
          flex: 1,
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        {renderGrid()}
        <div
          style={{
            marginLeft: 32,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            gap: 16,
          }}
        >
          <div
            style={{ display: "grid", gridTemplateColumns: "1fr 1fr", gap: 20 }}
          >
            <div
              style={{
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                gap: 4,
              }}
            >
              <Button size="large" onClick={() => Game.MoveLeft()}>
                ←
              </Button>
              <span style={{ fontSize: 12, color: "#666" }}>J</span>
            </div>
            <div
              style={{
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                gap: 4,
              }}
            >
              <Button size="large" onClick={() => Game.MoveRight()}>
                →
              </Button>
              <span style={{ fontSize: 12, color: "#666" }}>L</span>
            </div>
            <div
              style={{
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                gap: 2,
              }}
            >
              <Button size="large" onClick={() => Game.MagicLeft()}>
                ↙
              </Button>
              <span style={{ fontSize: 12, color: "#666" }}>A</span>
            </div>
            <div
              style={{
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                gap: 4,
              }}
            >
              <Button size="large" onClick={() => Game.MagicRight()}>
                ↘
              </Button>
              <span style={{ fontSize: 12, color: "#666" }}>D</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default IceMagic;
