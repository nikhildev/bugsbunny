import { DataTable, type TableColumn } from "@/components/DataTable";
import { getAllIssues } from "./api/api";
import { Badge } from "@/components/ui/badge";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { format, formatISO, parseISO } from "date-fns";

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

export const columns: TableColumn<Issue, any>[] = [
  {
    header: "Title",
    accessorKey: "title",
  },
  {
    header: "Created At",
    accessorKey: "created_at",
    cell: ({ row }) => {
      return (
        <>
          {format(
            parseISO(String(row.getValue("created_at"))),
            "dd.MM.yyyy HH:mm",
          )}
        </>
      );
    },
  },
  {
    header: "Updated At",
    accessorKey: "updated_at",
    cell: ({ row }) => {
      return (
        <>
          {format(
            parseISO(String(row.getValue("updated_at"))),
            "dd.MM.yyyy HH:mm",
          )}
        </>
      );
    },
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
  const queryClient = useQueryClient();
  const { data: issues } = useQuery({
    queryKey: ["issues"],
    queryFn: getAllIssues,
  });

  return (
    <div>
      <h1>Issues</h1>
      <DataTable columns={columns} data={(issues || []) as Issue[]} />
    </div>
  );
};
