import * as RadixTooltip from "@radix-ui/react-tooltip";
import "./Tooltip.scss";

interface TooltipProps {
  children: React.ReactNode;
  content: React.ReactNode;
  align?: "start" | "center" | "end";
}

const Tooltip: React.FC<TooltipProps> = ({
  children,
  content,
  align = "start",
}) => {
  return (
    <RadixTooltip.Provider delayDuration={200}>
      <RadixTooltip.Root>
        <RadixTooltip.Trigger asChild>
          <div>{children}</div>
        </RadixTooltip.Trigger>
        <RadixTooltip.Portal>
          <RadixTooltip.Content
            side="bottom"
            align={align}
            className="tooltip-content"
            sideOffset={4}
          >
            {content}
          </RadixTooltip.Content>
        </RadixTooltip.Portal>
      </RadixTooltip.Root>
    </RadixTooltip.Provider>
  );
};

export default Tooltip;
