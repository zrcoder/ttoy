import { Events } from "@wailsio/runtime";
import { Button, Select } from "antd";
import React, { useEffect, useRef, useState } from "react";
import { Game } from "../../bindings/github.com/zrcoder/ttoy/game/ballsort/index";
import type { Ball, Bottle, Level } from "../../bindings/github.com/zrcoder/ttoy/game/ballsort/models";
import { contentHeight } from "../components/common/layout";

const BallSort: React.FC = () => {
  const [bottles, setBottles] = useState<Bottle[]>([]);
  const [levels, setLevels] = useState<Level[]>([]);
  const [selectedLevel, setSelectedLevel] = useState(0);
  const [done, setDone] = useState(0);
  const [total, setTotal] = useState(0);
  const [cursor, setCursor] = useState(0);
  const [shiftBallIndex, setShiftBallIndex] = useState(-1);
  const [shiftBallColor, setShiftBallColor] = useState("");
  const bottleRefs = useRef<(HTMLDivElement | null)[]>([]);

  useEffect(() => {
    Game.Levels().then((l) => {
      setLevels(l || []);
      Game.SelectLevel(0);
    });

    const stop = Events.On(
      "ballsort:update",
      (ev: {
        data: {
          Bottles: Bottle[];
          DoneCount: number;
          TotalCount: number;
          Cursor: number;
          ShiftBallIndex: number;
          ShiftBallColor: string;
        };
      }) => {
        setBottles(ev.data.Bottles);
        setDone(ev.data.DoneCount);
        setTotal(ev.data.TotalCount);
        setCursor(ev.data.Cursor);
        setShiftBallIndex(ev.data.ShiftBallIndex);
        setShiftBallColor(ev.data.ShiftBallColor);
      },
    );
    return stop;
  }, []);

  const handleSelectBottle = async (index: number) => {
    await Game.SelectBottle(index);
  };

  const handleLevelChange = async (index: number) => {
    setSelectedLevel(index);
    await Game.SelectLevel(index);
  };

  const handleReset = async () => {
    await Game.Reset();
  };

  const onKey = (e: KeyboardEvent) => {
    const code = e.code;
    if (code.startsWith("Key") && code.length === 4) {
      const index = code.charCodeAt(3) - "A".charCodeAt(0);
      handleSelectBottle(index);
    }
  };

  useEffect(() => {
    window.addEventListener("keydown", onKey);
    return () => window.removeEventListener("keydown", onKey);
  }, []);

  const floatingBall =
    shiftBallIndex >= 0 && bottleRefs.current[shiftBallIndex]
      ? (() => {
          const ref = bottleRefs.current[shiftBallIndex]!;
          const rect = ref.getBoundingClientRect();
          const container = ref.parentElement!;
          const containerRect = container.getBoundingClientRect();
          return {
            left: rect.left - containerRect.left + (rect.width - 40) / 2,
            top: rect.top - containerRect.top - 44,
          };
        })()
      : null;

  const floatingBallStyle: React.CSSProperties | null =
    floatingBall && shiftBallColor != ""
      ? {
          position: "absolute",
          left: floatingBall.left,
          top: floatingBall.top,
          width: 40,
          height: 40,
          borderRadius: "50%",
          backgroundColor: shiftBallColor,
          boxShadow: "0 4px 12px rgba(0,0,0,0.3)",
          zIndex: 10,
          pointerEvents: "none",
        }
      : { display: "none" };

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
        <span>
          Done: {done}/{total}
        </span>
      </div>

      <div
        style={{
          height: "100%",
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
        }}
      >
        <div
          style={{
            display: "flex",
            gap: 16,
            flexWrap: "wrap",
            justifyContent: "center",
            position: "relative",
          }}
        >
          {floatingBallStyle && <div style={floatingBallStyle} />}
          {bottles.map((bottle, i) => (
            <div
              key={i}
              ref={(el) => {
                bottleRefs.current[i] = el;
              }}
              onClick={() => handleSelectBottle(i)}
              style={{
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                cursor: "pointer",
                padding: 8,
                border: `2px solid ${i === cursor ? "#1890ff" : "#ccc"}`,
                borderRadius: 8,
                minWidth: 60,
                boxShadow:
                  i === cursor ? "0 0 12px rgba(24, 144, 255, 0.5)" : "none",
              }}
            >
              <div
                style={{
                  height: 160,
                  display: "flex",
                  flexDirection: "column-reverse",
                  gap: 4,
                }}
              >
                {(bottle.Balls || []).map((ball: Ball | null, j: number) => (
                  <div
                    key={j}
                    style={{
                      width: 40,
                      height: 40,
                      borderRadius: "50%",
                      backgroundColor: ball?.Color ?? "transparent",
                    }}
                  />
                ))}
              </div>
              <span style={{ marginTop: 8, fontWeight: "bold" }}>
                {String.fromCharCode(65 + i)}
              </span>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default BallSort;
