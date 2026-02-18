import "./index.css";
import { Routes, Route } from "react-router";
import { Home } from "./pages/Home";
import { IssuesList } from "./pages/issues/IssuesList";
import { AppSidebar } from "./components/AppSidebar";
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from "./components/ui/sidebar";

export function App() {
  return (
    <SidebarProvider>
      <AppSidebar />
      <SidebarInset>
        {/* <SidebarTrigger /> */}
        <main className="flex-1">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/issues" element={<IssuesList />} />
          </Routes>
        </main>
      </SidebarInset>
    </SidebarProvider>
  );
}

export default App;
