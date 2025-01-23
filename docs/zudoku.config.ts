import type { ZudokuConfig } from "zudoku";

const config: ZudokuConfig = {
  basePath: "/",
  redirects: [
    {from: "/", to: "/overview"},
    {from: "/guides", to: "/guides/deployment"},
    {from: "/references", to: "/references/api"},
  ],
  page: {
    pageTitle: "",
    logoUrl: "/",
    logo: {
      src: {
        light: "/logo/outpost-logo-black.svg",
        dark: "/logo/outpost-logo-white.svg"
      },
      width: "110px",
    }
  },
  topNavigation: [
    { id: "docs", label: "Documentation", default: "overview" },
    // { id: "docs/api", label: "API Reference" },
    // { id: "/", label: "Website" },
  ],
  sidebar: {
    docs: [
      {
        type: "doc",
        label: "Overview",
        id: "overview",
      },
      {
        type: "doc",
        label: "Quickstart",
        id: "quickstart",
      },
      {
        type: "doc",
        label: "Concepts",
        id: "concepts",
      },
      {
        type: "doc",
        label: "Features",
        id: "features",
      },
      // {
      //   type: "doc",
      //   label: "Configuration",
      //   id: "docs/configuration",
      // },
      {
        type: "category",
        label: "Guides",
        collapsed: false,
        collapsible: false,
        items: [
          {
            type: "doc",
            label: "Deployment",
            id: "guides/deployment",
          },
          {
            type: "doc",
            label: "Dashboard Design",
            id: "guides/dashboard-design",
          },
        ],
      },
      {
        type: "category",
        label: "References",
        collapsed: false,
        collapsible: false,
        items: [
          {
            type: "doc",
            label: "Configuration",
            id: "references/configuration",
          },
          {
            type: "doc",
            label: "API",
            id: "references/api",
          },
        ],
      },
    ],
  },

  apis: {
    type: "file",
    input: "./apis/openapi.yaml",
    navigationId: "docs/api",
  },
  docs: {
    files: "/pages/**/*.{md,mdx}",
  },
};

export default config;
