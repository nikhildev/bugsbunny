import "./index.css";
import { Routes, Route } from "react-router";
import IssuesList from "./pages/issues/IssuesList";

export function App() {
  return (
    <Routes>
      <Route path="/" element={<IssuesList />} />
    </Routes>
  );
}

export default App;
