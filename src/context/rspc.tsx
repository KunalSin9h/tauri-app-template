import { QueryClient } from "@tanstack/react-query";
import { createClient } from "@rspc/client";
import { TauriTransport } from "@rspc/tauri";
import { createReactQueryHooks } from "@rspc/react";
import type { Procedures } from "../ts/bindings";

const client = createClient<Procedures>({
  transport: new TauriTransport(),
});

const queryClient = new QueryClient();

export const rspc = createReactQueryHooks<Procedures>();

export default function RspcProvider(props) {
  return (
    <rspc.Provider client={client} queryClient={queryClient}>
      {props.children}
    </rspc.Provider>
  );
}
