import type { Issue } from "@/models/issue";

export const getAllIssues = async (): Promise<Issue[] | undefined> => {
  const res = await fetch("http://localhost:8080/issues");
  if (!res.ok) {
    return undefined;
  }
  const data = await res.json();
  return data as Issue[] | undefined;
};
