export type MethodCall = {
  id: number;
  method: string;
  params: unknown[];
};
export type MethodResult<T> = {
  id: number;
  ok: T | null;
  err: string | null;
};
