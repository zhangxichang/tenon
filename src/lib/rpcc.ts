import type { MethodCall, MethodResult, Task } from "./types";

export class RPConnection {
	private websocket: WebSocket;

	private constructor(websocket: WebSocket) {
		this.websocket = websocket;
	}
	static async new() {
		return new RPConnection(
			await new Promise((resolve, reject) => {
				const websocket = new WebSocket("ws://localhost:10270/ws");
				websocket.onopen = () => resolve(websocket);
				websocket.onerror = () => reject(new Error("连接失败"));
			}),
		);
	}
	cleanup() {
		this.websocket.close();
	}
	async call<T>(method: string, params?: unknown[]) {
		return await new Promise<T>((resolve, reject) => {
			this.websocket.onerror = () => reject(new Error("连接失败"));
			this.websocket.send(
				JSON.stringify({
					id: 0,
					method,
					params: params ?? [],
				} satisfies MethodCall),
			);
			this.websocket.onmessage = (e: MessageEvent<string>) => {
				const { ok, err }: MethodResult<T> = JSON.parse(e.data);
				if (err !== null) {
					reject(err);
				}
				resolve(ok);
			};
		});
	}
	async get_all_task() {
		return await this.call<Task[]>("get_all_task");
	}
}
