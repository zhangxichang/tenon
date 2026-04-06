import { RPConnection } from "~/lib/rpcc";

export class ViewportStore {
	rpcc: RPConnection;

	private constructor(rpcc: RPConnection) {
		this.rpcc = rpcc;
	}
	static async new() {
		return new ViewportStore(await RPConnection.new());
	}
	cleanup() {
		this.rpcc.cleanup();
	}
}
