"use client";

import { withObservables } from '@nozbe/watermelondb/react'
import { useGetCommands } from './hooks/useQuery';
import { CreateCommandButton, CreateIntentButton, SyncButton } from "./buttons";
import IntentsList from './intents';
import { useGetIntents } from './hooks/useQuery';
import { Command } from '@dilocash/database/local/model/commmand';

const Commands = () => {
  const commands = useGetCommands();
  return <div>
      <div style={{ marginBottom: 20 }}>
        <SyncButton />
      </div>

      <h1>Commands:</h1>

      <div>
        <CreateCommandButton />
      </div>

      <EnhancedCommandsList commands={commands} />
    </div>
}

export default Commands;

const CommandsList = ({ commands }: { commands: Command[] }) => (
  <div>
    {commands.map((command, i) => (
      <CommandItem key={i} command={command} />
    ))}
  </div>
);

const EnhancedCommandsList = withObservables(["commands"], ({ commands }) => ({
  commands,
}))(CommandsList);

const CommandItem = ({ command }: { command: Command }) => {
  const intents = useGetIntents(command.id); // we need to use this hook for observability to work.

  return (
    <div>
      <h1 style={{ fontSize: 20, margin: 15 }}>{command.status}</h1>
      <button style={{ color: "red" }} onClick={() => command.delete()}>
        <p>{command.status}</p>
      </button>
      <div>
        <CreateIntentButton command={command} />
        <IntentsList intents={intents} />
      </div>
    </div>
  );
};