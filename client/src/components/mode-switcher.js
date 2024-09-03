"use client";

import { Switch } from "@/components/ui/switch";
import { Label } from "@/components/ui/label";
import { useTheme } from "@/app/ThemeProvider";
import { cn } from "@/lib/utils";

import { MoonIcon as Moon, SunMediumIcon as Sun } from "lucide-react";

const ModeSwitcher = (props) => {
  const { theme, toggleTheme } = useTheme();
  return (
    <div className={cn(props.className, "flex items-center space-x-2")}>
      <Switch id="mode-switch" onClick={toggleTheme} width={20} height={10}>
        {theme === "dark" ? <Sun /> : <Moon />}
      </Switch>
      <Label className="text-sm text-text-muted" htmlFor="mode-switch">
        {theme === "dark" ? "Dark Mode" : "Light Mode"}
      </Label>
    </div>
  );
};

export default ModeSwitcher;
