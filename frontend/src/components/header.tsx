import { Save, SquarePen } from "lucide-react";
import { ModeToggle } from "./mode-toggle";
import { Button } from "./ui/button";

function Header() {
  return (
    <nav className="h-12 bg-background border flex items-center justify-between w-screen py-2 px-4">
      <div className="flex items-center gap-1 ">
        <Button variant={"outline"}>
          <SquarePen />
        </Button>
        <div className="relative inline-block">
          <input
            name="filename"
            type="text"
            maxLength={55}
            autoComplete="off"
            value="Untitled"
            className="box-content text-[1.3em] outline-none focus:border-transparent focus:ring-0"
          />
        </div>
      </div>

      <div className="flex gap-2">
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
    </nav>
  );
}

export default Header;
