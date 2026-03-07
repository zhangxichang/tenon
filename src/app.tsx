import "~/app.css";
import { Router } from "@solidjs/router";
import { FileRoutes } from "@solidjs/start/router";
import { Shell } from "./shell";

export default function App() {
  return (
    <Router root={Shell}>
      <FileRoutes />
    </Router>
  );
}
