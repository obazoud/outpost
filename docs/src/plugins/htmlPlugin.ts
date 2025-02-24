import React from "react";
import type { ZudokuPlugin } from "zudoku";

export const htmlPlugin = ({
  headScript,
}: {
  headScript: string;
}): ZudokuPlugin => {
  return {
    getHead: () => {
      return React.createElement("script", null, headScript);
    },
  };
};
