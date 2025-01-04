import { FC } from "react";
import "./Badge.scss";

interface BadgeProps {
  success?: boolean;
  danger?: boolean;
  primary?: boolean;
  text: string | number;
}

const Badge: FC<BadgeProps> = ({
  text,
  success = false,
  danger = false,
  primary = false,
}) => {
  const className = `badge${success ? " badge__success" : ""}${
    danger ? " badge__danger" : ""
  }${primary ? " badge__primary" : ""}`;

  return <span className={className}>{text}</span>;
};

export default Badge;
