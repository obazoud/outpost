import { Link } from "react-router-dom";
import { FC, PropsWithChildren } from "react";
import "./Button.scss";

interface ButtonProps {
  primary?: boolean;
  to?: string;
  onClick?: () => void;
  disabled?: boolean;
  type?: "button" | "submit" | "reset";
  className?: string;
}

const Button: FC<PropsWithChildren<ButtonProps>> = ({
  primary = false,
  to,
  onClick,
  disabled = false,
  type = "button",
  children,
  className,
}) => {
  className = `button${primary ? " button__primary" : ""}${
    disabled ? " button__disabled" : ""
  } ${className}`;

  if (to) {
    return (
      <Link
        to={to}
        className={className}
        {...(disabled && { onClick: (e) => e.preventDefault() })}
      >
        {children}
      </Link>
    );
  }

  return (
    <button
      onClick={onClick}
      className={className}
      disabled={disabled}
      type={type}
    >
      {children}
    </button>
  );
};

export default Button;
