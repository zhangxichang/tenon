import { useIsRouting, type RouteSectionProps } from "@solidjs/router";
import { ShellStore } from "./stores/shell";
import { ErrorBoundary, For, onCleanup, Show, Suspense } from "solid-js";
import { twMerge } from "tailwind-merge";
import { ShellContext } from "./components/context";
import Error from "./components/widgets/error";
import Loading from "./components/widgets/loading";

export function Shell(props: RouteSectionProps) {
  const store = ShellStore.new();
  const is_routing = useIsRouting();
  onCleanup(() => store.cleanup());
  return (
    <ShellContext.Provider value={store}>
      <div class="absolute w-dvw h-dvh flex flex-col bg-base-200">
        <div class="flex-1 flex relative min-h-0">
          <Show when={is_routing()}>
            <progress class="progress progress-primary absolute rounded-none h-0.5" />
          </Show>
          <ErrorBoundary fallback={(error) => <Error error={error as Error} />}>
            <Suspense fallback={<Loading />}>{props.children}</Suspense>
          </ErrorBoundary>
        </div>
        <div class="toast">
          <For each={store.toaster.items}>
            {(v) => {
              const level_style = {
                info: "alert-info",
                success: "alert-success",
                warning: "alert-warning",
                error: "alert-error",
              }[v.level];
              return (
                <div class={`${twMerge("alert", level_style)}`}>
                  {v.content}
                </div>
              );
            }}
          </For>
        </div>
      </div>
    </ShellContext.Provider>
  );
}
