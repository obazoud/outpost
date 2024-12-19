import { Link } from "react-router-dom";
import { FC, PropsWithChildren } from "react";
import "./Button.scss";
import { Loading } from "../Icons";

interface ButtonProps {
  to?: string;
  onClick?: () => void;
  disabled?: boolean;
  type?: "button" | "submit" | "reset";
  className?: string;
  loading?: boolean;
  primary?: boolean;
  danger?: boolean;
}

const Button: FC<PropsWithChildren<ButtonProps>> = ({
  primary = false,
  to,
  onClick,
  disabled = false,
  type = "button",
  children,
  className,
  loading = false,
  danger = false,
}) => {
  className = `button${primary ? " button__primary" : ""}${
    disabled || loading ? " button__disabled" : ""
  } ${loading ? " button__loading" : ""} ${danger ? " button__danger" : ""} ${
    className || ""
  }`;

  if (to) {
    return (
      <Link
        to={to}
        className={className}
        {...(disabled && { onClick: (e) => e.preventDefault() })}
      >
      <span>{children}</span>
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
      <span>{children}</span>
      {loading && (
        <div className="button__loading-container">
          <Loading />
        </div>
      )}
    </button>
  );
};

export default Button;
