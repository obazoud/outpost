import { Link } from "react-router-dom";
import { FC, PropsWithChildren } from "react";
import "./Button.scss";

interface ButtonProps {
  primary?: boolean;
  to?: string;
  onClick?: () => void;
}

const Button: FC<PropsWithChildren<ButtonProps>> = ({
  primary = false,
  to,
  onClick,
  children,
}) => {
  const className = `button${primary ? " button__primary" : ""}`;

  if (to) {
    return (
      <Link to={to} className={className}>
        {children}
      </Link>
    );
  }

  return (
    <button onClick={onClick} className={className}>
      {children}
    </button>
  );
};

export default Button;
