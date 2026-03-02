import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupAction,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar";
import { Link } from "react-router";
import { HomeIcon, ListIcon, ComponentIcon } from "lucide-react";

export const AppSidebar = () => {
  return (
    <Sidebar>
      <SidebarHeader>
        <h1 className="text-2xl font-semibold text-gray-700">Bugs Bunny</h1>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          {/* <SidebarGroupLabel>Pages</SidebarGroupLabel> */}
          <SidebarGroupContent>
            <SidebarMenu>
              <SidebarMenuItem>
                <SidebarMenuButton>
                  <Link to="/" className="flex items-center gap-2">
                    <HomeIcon className="size-4" />
                    Home
                  </Link>
                </SidebarMenuButton>
              </SidebarMenuItem>
            </SidebarMenu>
            <SidebarMenu>
              <SidebarMenuItem>
                <SidebarMenuButton>
                  <Link to="/issues" className="flex items-center gap-2">
                    <ListIcon className="size-4" />
                    Issues List
                  </Link>
                </SidebarMenuButton>
              </SidebarMenuItem>
            </SidebarMenu>
            <SidebarMenu>
              <SidebarMenuItem>
                <SidebarMenuButton>
                  <Link to="/projects" className="flex items-center gap-2">
                    <ComponentIcon className="size-4" />
                    Projects
                  </Link>
                </SidebarMenuButton>
              </SidebarMenuItem>
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter />
    </Sidebar>
  );
};
