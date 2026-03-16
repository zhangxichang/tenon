export type MethodCall = {
  id: number;
  method: string;
  params: unknown[];
};
export type MethodResult<T = null> = {
  id: number;
  ok: T;
  err: string | null;
};
