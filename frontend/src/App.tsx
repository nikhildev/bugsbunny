import "./index.css";
import { Routes, Route } from "react-router";
import { Home } from "./pages/Home";
import {IssuesList} from "./pages/issues/IssuesList";

export function App() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/issues" element={<IssuesList />} />
    </Routes>
  );
}

export default App;
