import { Button, Select } from "antd";
import React, { useEffect, useState } from "react";
import { contentHeight } from "../components/common/layout";
import { Events } from "@wailsio/runtime";
import * as Game from "../../bindings/github.com/zrcoder/ttoy/game/hanoi/game";
import type {
  Level,
  PileInfo,
} from "../../bindings/github.com/zrcoder/ttoy/game/hanoi/models";

const diskColors = ["red", "green", "blue", "yellow", "brown", "pink"];

const Hanoi: React.FC = () => {
  const [piles, setPiles] = useState<PileInfo[]>([]);
  const [levels, setLevels] = useState<Level[]>([]);
  const [selectedLevel, setSelectedLevel] = useState(0);
  const [steps, setSteps] = useState(0);

  useEffect(() => {
    Game.Levels().then((l) => {
      setLevels(l || []);
      Game.SelectLevel(0).then(() => {
        Game.Piles().then((p) => setPiles(p || []));
      });
    });

    const stop = Events.On(
      "hanoi:update",
      (ev: { data: { Piles: PileInfo[]; Steps: number } }) => {
        setPiles(ev.data.Piles);
        setSteps(ev.data.Steps);
      },
    );
    return stop;
  }, []);

  const handleSelectPile = async (index: number) => {
    await Game.SelectPileByIndex(index);
  };

  const handleLevelChange = async (index: number) => {
    setSelectedLevel(index);
    await Game.SelectLevel(index);
  };

  const handleReset = async () => {
    await Game.Reset();
  };

  return (
    <div
      style={{
        padding: 20,
        height: contentHeight,
        position: "relative",
      }}
    >
      <div
        style={{
          position: "absolute",
          top: 20,
          left: 20,
          display: "flex",
          gap: 16,
          alignItems: "center",
        }}
      >
        <Select
          value={selectedLevel}
          onChange={handleLevelChange}
          options={levels.map((l, i) => ({ label: l.Label, value: i }))}
          style={{ width: 150 }}
        />
        <Button onClick={handleReset}>Reset</Button>
        <span>Steps: {steps}</span>
      </div>

      <div
        style={{
          height: "100%",
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
        }}
      >
        <div style={{ display: "flex", gap: 32, justifyContent: "center" }}>
          {piles.map((pile, i) => (
            <div
              key={i}
              onClick={() => handleSelectPile(i)}
              style={{
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                cursor: "pointer",
                padding: 16,
                border: "1px solid #ccc",
                borderRadius: 8,
                minWidth: 120,
              }}
            >
              <div
                style={{
                  height: 200,
                  display: "flex",
                  flexDirection: "column-reverse",
                  gap: 2,
                  justifyContent: "flex-start",
                }}
              >
                {(pile.Disks || []).map((id, j) => (
                  <div
                    key={j}
                    style={{
                      width: id * 15 + 30,
                      height: 25,
                      borderRadius: 4,
                      backgroundColor:
                        diskColors[id % diskColors.length] || "gray",
                    }}
                  />
                ))}
              </div>
              <span style={{ marginTop: 8, fontWeight: "bold", fontSize: 18 }}>
                {String.fromCharCode(65 + i)}
              </span>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Hanoi;
