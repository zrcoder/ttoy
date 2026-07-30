let audioCtx: AudioContext | null = null;
let soundEnabled = true;

export const isSoundEnabled = () => soundEnabled;
export const setSoundEnabled = (enabled: boolean) => {
  soundEnabled = enabled;
};

const getAudioContext = () => {
  if (!audioCtx) {
    audioCtx = new AudioContext();
  }
  return audioCtx;
};

export const playMoveSound = () => {
  if (!soundEnabled) return;
  const ctx = getAudioContext();
  const osc = ctx.createOscillator();
  const gain = ctx.createGain();

  osc.connect(gain);
  gain.connect(ctx.destination);

  osc.type = "sine";
  osc.frequency.setValueAtTime(440, ctx.currentTime);
  osc.frequency.exponentialRampToValueAtTime(880, ctx.currentTime + 0.05);

  gain.gain.setValueAtTime(0.3, ctx.currentTime);
  gain.gain.exponentialRampToValueAtTime(0.01, ctx.currentTime + 0.08);

  osc.start(ctx.currentTime);
  osc.stop(ctx.currentTime + 0.08);
};

export const playMagicSound = () => {
  if (!soundEnabled) return;
  const ctx = getAudioContext();

  const osc1 = ctx.createOscillator();
  const osc2 = ctx.createOscillator();
  const gain = ctx.createGain();

  osc1.connect(gain);
  osc2.connect(gain);
  gain.connect(ctx.destination);

  osc1.type = "triangle";
  osc2.type = "sine";

  osc1.frequency.setValueAtTime(523, ctx.currentTime);
  osc1.frequency.exponentialRampToValueAtTime(1047, ctx.currentTime + 0.1);

  osc2.frequency.setValueAtTime(659, ctx.currentTime);
  osc2.frequency.exponentialRampToValueAtTime(1319, ctx.currentTime + 0.15);

  gain.gain.setValueAtTime(0.2, ctx.currentTime);
  gain.gain.exponentialRampToValueAtTime(0.01, ctx.currentTime + 0.2);

  osc1.start(ctx.currentTime);
  osc1.stop(ctx.currentTime + 0.2);
  osc2.start(ctx.currentTime);
  osc2.stop(ctx.currentTime + 0.2);
};
