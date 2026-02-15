import { DataTable } from "@/components/DataTable";
import { getAllIssues } from "./api/api";
import { useState, useEffect } from "react";
import type { ColumnDef } from "@tanstack/react-table";
import { Badge } from "@/components/ui/badge";

interface Issue {
  id: string;
  title: string;
  description: string;
  status: string;
  createdAt: string;
  updatedAt: string;
  type: string;
  reporter: string;
  componentId: string;
  priority: string;
  severity: string;
}

export const columns: ColumnDef<Issue, any>[] = [
  {
    header: "Title",
    accessorKey: "title",
  },
  {
    header: "Created At",
    accessorKey: "created_at",
  },
  {
    header: "Updated At",
    accessorKey: "updated_at",
  },
  {
    header: "Type",
    accessorKey: "type",
    cell: ({ row }) => {
      return <>{row.original.type.toUpperCase()}</>;
    },
  },
  {
    header: "Reporter",
    accessorKey: "reporter",
  },
  {
    header: "Component ID",
    accessorKey: "component_id",
  },
  {
    header: "Priority",
    accessorKey: "priority",
    cell: ({ row }) => {
      return <Badge>{row.original.priority.toUpperCase()}</Badge>;
    },
  },
  {
    header: "Severity",
    accessorKey: "severity",
    cell: ({ row }) => {
      return <Badge>{row.original.severity.toUpperCase()}</Badge>;
    },
  },
];

export const IssuesList = () => {
  const [issues, setIssues] = useState<Issue[]>([]);
  useEffect(() => {
    getAllIssues().then(setIssues);
  }, []);

  return <div>
    <h1>Issues</h1>
    <DataTable columns={columns} data={issues} />
  </div>;
};