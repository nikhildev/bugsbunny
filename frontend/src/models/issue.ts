export interface Issue {
  id: string;
  title: string;
  description: string;
  status: IssueStatus;
  created_at: string;
  updated_at: string;
  type: IssueType;
  reporter: User;
  component_id: string;
  priority: Priority;
  severity: Severity;
}

export type IssueStatus =
  | "new"
  | "resolved"
  | "in_progress"
  | "reopened"
  | "blocked"
  | "on_hold";
export type IssueType =
  | "bug"
  | "feature"
  | "support"
  | "improvement"
  | "documentation";
export type Priority = "low" | "medium" | "high" | "critical";
export type Severity = "low" | "medium" | "high" | "critical";

export interface User {
  id: string;
  name: string;
  email: string;
}
