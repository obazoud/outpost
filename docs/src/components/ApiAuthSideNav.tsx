export const ApiAuthSideNav = ({
  location,
}: {
  location: { pathname: string };
}) => {
  if (location.pathname.startsWith("/api/") === false) {
    return null;
  }
  return (
    <a
      type="button"
      aria-controls="radix-:r4:"
      aria-expanded="true"
      data-state="open"
      className={`group flex items-center gap-2 px-[--padding-nav-item] py-1.5 rounded-lg hover:bg-accent ${
        location.pathname === "/api/authentication"
          ? "text-primary"
          : "text-foreground/80"
      } text-start font-medium cursor-pointer ${
        location.pathname === "/api/authentication" ? "active" : ""
      }`}
      href="/docs/api/authentication"
      data-discover="true"
      data-active={location.pathname === "/api/authentication"}
    >
      <div className="flex items-center gap-2 justify-between w-full">
        <div className="truncate">Authentication</div>
      </div>
    </a>
  );
};
