import { Save, SquarePen } from "lucide-react";
import { ModeToggle } from "./mode-toggle";
import { Button } from "./ui/button";

function Header() {
  return (
    <header className="h-12 bg-background border flex items-center justify-between w-screen py-2 px-4">
      <div className="flex items-center gap-1 ">
        <Button variant={"outline"} className="cursor-pointer">
          <SquarePen />
        </Button>
        <div className="w-25 sm:w-auto ">
          <p className="font-medium truncate">Untitled</p>
        </div>
      </div>

      <div className="flex gap-1">
        <Button size="default" className="cursor-pointer" variant={"outline"}>
          <Save />
          Save
        </Button>
        <Button size="default" className="cursor-pointer">
          Login
        </Button>
        <Button size="default" className="cursor-pointer" variant={"secondary"}>
          Sign Up
        </Button>

        <ModeToggle />
      </div>
    </header>
  );
}

export default Header;
