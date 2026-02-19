import { DataTable, type TableColumn } from "@/components/DataTable";
import { getAllIssues } from "./api/api";
import { Badge } from "@/components/ui/badge";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { getLocalisedTimestamp } from "@/lib/datetime";
import type { Issue } from "@/models/issue";

export const columns: TableColumn<Issue, any>[] = [
  {
    header: "Title",
    accessorKey: "title",
  },
  {
    header: "Created At",
    accessorKey: "created_at",
    cell: ({ row }) => {
      return <>{getLocalisedTimestamp(row.getValue("created_at") as string)}</>;
    },
  },
  {
    header: "Updated At",
    accessorKey: "updated_at",
    cell: ({ row }) => {
      return <>{getLocalisedTimestamp(row.getValue("updated_at") as string)}</>;
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
  const { data: issues, isLoading } = useQuery({
    queryKey: ["issues"],
    queryFn: getAllIssues,
  });

  return (
    <div>
      <h1 className="text-2xl font-bold mb-4">Issues</h1>

      <DataTable columns={columns} data={(issues || []) as Issue[]} />
    </div>
  );
};
