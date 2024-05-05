import * as React from 'react';

export const BreakPoints = {
  Mobile: 768,
};

export const MediaQueries = {
  Mobile: `(max-width: ${BreakPoints.Mobile})`,
};

export const useMediaQueries = () => {
  const getMatches = () => {
    return {
      mobile:
        typeof window !== 'undefined'
          ? window.matchMedia(MediaQueries.Mobile).matches
          : false,
    };
  };

  const [matches] = React.useState(getMatches);
  const initialMatchesRef = React.useRef(matches);

  return {
    matches,
    initialMatches: initialMatchesRef.current,
  };
};
