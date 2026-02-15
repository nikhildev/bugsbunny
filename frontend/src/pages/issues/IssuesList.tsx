import { getAllIssues } from "./api/api";
import { useState, useEffect } from "react";

interface Issue {
  id: string;
  title: string;
  description: string;
  status: string;
}

export const IssuesList = () => {
  const [issues, setIssues] = useState<Issue[]>([]);
  useEffect(() => {
    getAllIssues().then(setIssues);
  }, []);

  return (
  <div>
    <h1>Issues</h1>
    <ul>
      {issues.map((issue) => (
        <li key={issue.id}>{issue.title}</li>
        ))}
        </ul>
    </div>
  );
};