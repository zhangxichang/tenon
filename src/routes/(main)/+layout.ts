import { MainStore } from "~/stores/main";
import type { LayoutLoad } from "./$types";

export const load: LayoutLoad = async () => {
  return {
    main_store: await MainStore.new(),
  };
};
