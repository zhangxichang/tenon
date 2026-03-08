import type { Store } from "./interface";

export class MainStore implements Store {
  websocket: WebSocket;

  private constructor(websocket: WebSocket) {
    this.websocket = websocket;
  }
  static async new() {
    return new MainStore(
      await new Promise((resolve, reject) => {
        const websocket = new WebSocket("ws://localhost:10270/ws");
        websocket.onopen = () => resolve(websocket);
        websocket.onerror = () => reject(new Error("连接失败"));
      }),
    );
  }
  async cleanup() {
    this.websocket.close();
  }
}
