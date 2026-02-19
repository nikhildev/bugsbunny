export interface Issue {
  id: string;
  title: string;
  description: string;
  status: string;
  created_at: string;
  updated_at: string;
  type: string;
  reporter: string;
  component_id: string;
  priority: string;
  severity: string;
}
