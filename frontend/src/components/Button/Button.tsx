import { type PropsWithChildren } from "react";
import styles from "./Button.module.scss";

interface ButtonProps {
  variant: "primary" | "secondary" | "tertiary" | "danger" | "ghost";
  onClick: () => void;
  icon?: string;
  className?: string;
  loading?: boolean;
  disabled?: boolean;
  type?: "button" | "submit" | "reset";
}

function Button({
  variant,
  className,
  onClick,
  icon,
  disabled,
  type = "button",
  loading,
  children,
}: PropsWithChildren<ButtonProps>) {
  return (
    <button
      type={type}
      disabled={disabled || loading}
      className={[styles.button, styles[variant], className].filter(Boolean).join(" ")}
      onClick={onClick}
    >
      {loading && <span>Loading...</span>}
      {icon && <span>{icon}</span>}
      {children}
    </button>
  );
}

export default Button;
