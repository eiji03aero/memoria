import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import LanguageDetector from 'i18next-browser-languagedetector';

const Languages = ['ja', 'en'] as const;
type Language = (typeof Languages)[number];
const DefaultLanguage = 'en';

export const saveLocaleLanguage = (lang: Language) => {
  window.localStorage.setItem('mmr-lang', lang);
};

export const getLocaleLanguage = (): Language | undefined => {
  const lang = window.localStorage.getItem('mmr-lang');
  if (!lang || !['ja', 'en'].includes(lang)) {
    return undefined;
  }

  return lang as Language;
};

export const changeLocaleLanguage = (lang: Language) => {
  i18n.changeLanguage(lang);
  saveLocaleLanguage(lang);
};

i18n
  .use(initReactI18next)
  .use(LanguageDetector)
  .init({
    fallbackLng: DefaultLanguage,
    resources: {
      en: {
        translation: require('./locales/en.json'),
      },
      ja: {
        translation: require('./locales/ja.json'),
      },
    },
  });

const initialLanguage = getLocaleLanguage();
if (initialLanguage) {
  i18n.changeLanguage(initialLanguage);
}

export default i18n;
