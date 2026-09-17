export function GET() {
  return new Response(
    JSON.stringify({
      $schema: "https://json-schema.org/draft/2020-12/schema",
      $id: "entry.schema.json",
      title: "Valid Keybinding Entry",
      description: "A keybinding to run a shell command in a view.",
      type: "object",
      required: ["key"],
      properties: {
        key: {
          title: "Bound Key",
          description: "The combination of keys that trigger the command.",
          type: "string",
        },
        name: {
          title: "Command name",
          description: "A descriptive name for the command",
          type: "string",
        },
        command: {
          title: "Bound Command",
          description:
            "The shell command that runs when you press the key combination.",
          type: "string",
        },
        successMessage: {
          title: "Success message",
          description:
            "A short footer message shown after the custom command exits successfully.",
          type: "string",
        },
        footer: {
          title: "Footer action",
          description:
            'Set to "agent" to expose this command as the clickable agent control in PR and issue views.',
          type: "string",
          enum: ["agent"],
        },
        builtin: {
          title: "Builtin Command",
          description:
            "One of gh-dash's builtin commands that will run when you press the key combination",
          type: "string",
        },
      },
    }),
  );
}
