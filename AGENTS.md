# Repository agent guidance

This repository keeps a generated code knowledge graph in `graphify-out/`.

- Before broad codebase exploration, run `graphify query "<question>"`.
- Use `graphify path "<A>" "<B>"` for call or dependency paths.
- Use `graphify explain "<concept>"` for focused architecture context.
- After changing code, run `graphify update .` before committing.
- Never add runtime secrets, Firebase credentials, `google-services.json`, database dumps, or production environment files to the graph or repository.
