/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

"use client";

import { Transport } from "@connectrpc/connect";
import CommandsBar from "./commands-bar";
import CommandsListView from "./commands-list";
import { VStack } from "../ui/vstack";
import { isWeb } from "@gluestack-ui/utils/nativewind-utils";
import TopBar from "./top-bar";

const CommandsView = ({ transport, className }: { transport: Transport, className?: string }) => {
  return (
    <VStack className={`${className}`}>
      <TopBar />
      <CommandsListView className="flex-1" />
      <CommandsBar transport={transport} className={`px-2 pt-2 ${isWeb ? "mb-5" : ""}`} />
    </VStack>
  );
};

export default CommandsView;
