import * as DropdownMenu from "@radix-ui/react-dropdown-menu";
import { FC, PropsWithChildren } from "react";
import Button from "../Button/Button";
import "./Dropdown.scss";
import Badge from "../Badge/Badge";

interface DropdownProps {
  trigger_icon?: React.ReactNode;
  trigger: React.ReactNode;
  disabled?: boolean;
  badge_count?: number;
}

const Dropdown: FC<PropsWithChildren<DropdownProps>> = ({
  trigger_icon,
  trigger,
  children,
  badge_count,
  disabled = false,
}) => {
  return (
    <DropdownMenu.Root>
      <DropdownMenu.Trigger asChild disabled={disabled}>
        <div>
          <Button>
            {trigger_icon}
            {trigger}
            {badge_count && badge_count > 0 ? (
              <Badge text={String(badge_count)} />
            ) : null}
          </Button>
        </div>
      </DropdownMenu.Trigger>
      <DropdownMenu.Portal>
        <DropdownMenu.Content
          className="dropdown-content"
          sideOffset={4}
          align="start"
        >
          {children}
        </DropdownMenu.Content>
      </DropdownMenu.Portal>
    </DropdownMenu.Root>
  );
};

export default Dropdown;
