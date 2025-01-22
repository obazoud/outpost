import CONFIGS from "../config";

const getLogo = () => {
  const logo =
    document.body.dataset.theme === "dark" && CONFIGS.LOGO_DARK
      ? CONFIGS.LOGO_DARK
      : CONFIGS.LOGO;

  return logo;
};

export default getLogo;
