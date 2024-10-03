import { cn } from "@/lib/utils";
import {
  FilePlus,
  LayoutDashboard,
  ChartNoAxesCombined,
  FolderUp,
  FolderDown,
} from "lucide-react";

export default function QuickActionsPanel(props) {
  return (
    <div className={cn("w-full h-full ", props.className)}>
      <h2 className="px-4 mb-8 text-2xl text-text">Get started</h2>
      <div className="grid items-center grid-cols-4 gap-2 px-4 justify-items-start">
        <div className="w-full h-full p-2 rounded-xl bg-secondary text-text-muted">
          <div className="flex flex-col justify-around w-full h-full gap-6 p-4 tracking-wide border-2 border-dashed rounded-sm border-text-muted/80 hover:text-text/70 hover:border-text/70">
            <div className="space-x-2">
              <span>Create New Dashboard</span>
              <span className="px-1 ml-1 font-mono text-xs tracking-wider border-2 w-fit text-text border-gray-600/70">
                SHIFT + D
              </span>
            </div>
            <LayoutDashboard
              height={90}
              width={90}
              className="mx-auto my-auto stroke-text-muted/40"
              strokeWidth={1}
            />
          </div>
        </div>
        <div className="w-full h-full p-2 rounded-xl bg-secondary text-text-muted">
          <div className="flex flex-col justify-around w-full h-full gap-6 p-4 tracking-wide border-2 border-dashed rounded-sm border-text-muted/80 hover:text-text/70 hover:border-text/70">
            <div className="space-x-2">
              <span>Generate Chart</span>
              <span className="px-1 ml-1 font-mono text-xs tracking-wider border-2 w-fit text-text border-gray-600/70">
                SHIFT + C
              </span>
            </div>
            <ChartNoAxesCombined
              height={100}
              width={100}
              className="mx-auto my-auto stroke-text-muted/40"
              strokeWidth={1}
            />
          </div>
        </div>
        <div className="w-full h-full p-2 rounded-xl bg-secondary text-text-muted">
          <div className="flex flex-col justify-around w-full h-full gap-6 p-4 tracking-wide border-2 border-dashed rounded-sm border-text-muted/80 hover:text-text/70 hover:border-text/70">
            <div className="space-x-2">
              <span>Start New Project</span>
              <span className="px-1 ml-1 font-mono text-xs tracking-wider border-2 w-fit text-text border-gray-600/70">
                SHIFT + P
              </span>
            </div>
            <FilePlus
              height={100}
              width={100}
              className="mx-auto my-auto stroke-text-muted/40"
              strokeWidth={1}
            />
          </div>
        </div>
        <div className="w-full h-full p-2 rounded-xl bg-secondary text-text-muted">
          <div className="flex flex-col justify-around w-full h-full gap-6 p-4 tracking-wide border-2 border-dashed rounded-sm border-text-muted/80 hover:text-text/70 hover:border-text/70">
            <div className="space-x-2">
              <span>Import</span>
              <span className="px-1 ml-1 font-mono text-xs tracking-wider border-2 w-fit text-text border-gray-600/70">
                SHIFT + I
              </span>
            </div>
            <FolderDown
              height={100}
              width={100}
              className="mx-auto my-auto stroke-text-muted/40"
              strokeWidth={1}
            />
          </div>
        </div>
      </div>
    </div>
  );
}
