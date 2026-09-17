# Devlog - 2026-09-18

### `just` kept breaking my local agent, so I built `wallop`

(reformatted by ai)

I run Hermes on Qwen3 35B on Windows, 12GB VRAM, hard 65k context. No cloud.

My Python scripts work fine in terminal. Ask the agent to run them: round 1 ok, round 3 it starts mangling `C:\Users\...\demo\system-sentinel\sentinel.py`, round 5 it hallucinates args and dies.

Token math: a full Windows path + payload is ~180 tokens. Repeated for 5 rounds = 1k+ tokens wasted just on paths. At 65k, that's my system prompt.

I tried `just`. For me, it's great. For the agent, worse.

1.  `justfile` has to be read every time to list recipes - another 1k tokens.
2.  `just run a check` is valid English. The model reads `just` as an adverb, not a tool. It starts narrating instead of calling.
3.  Recipe still contains the long path inside.

So I built the dumbest fix: an alias table that lives outside context. Here is the demo. And I mark this done for my future references.

**wallop**

- `wallop register <cmd>` -> saves full command to `commands.txt`
- `wallop check` -> lists as `+1, +2, +3` (~30 tokens)
- `wallop run +1` -> exec. 8 tokens, immutable. Agent no needs to see the path.

`just` = human alias. `wallop` = agent alias. The `+N` syntax is intentionally non-prose - model can't mistake it for conversation.

Same logic, but my 5-round task now saves ~1.5k tokens. Enough to keep the agent alive.

It's a Go binary, macOS `zsh` in README because I scaffolded fast, but I use it on Windows daily. No sandbox, not secure, just a toy that stopped the crazy.

If you run 35B on 12GB, you know this pain.