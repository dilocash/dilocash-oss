import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import LanguageDetector from 'i18next-browser-languagedetector'; // Web only
import es from './locales/es.json';
import en from './locales/en.json';

export const initI18n = (isWeb: boolean, extraLng?: string) => {
  const instance = i18n.use(initReactI18next);

  // Only use the automatic detector if we are on the web
  if (isWeb) {
    instance.use(LanguageDetector);
  }

  return instance.init({
    resources: {
      es: { translation: es },
      en: { translation: en },
    },
    // IMPORTANT: Don't define 'lng' here for the detector to work
    fallbackLng: 'en', 
    supportedLngs: ['es', 'en'],
    // Detector configuration
    detection: {
      order: ['querystring', 'navigator', 'htmlTag']
    },
    interpolation: {
      escapeValue: false,
    },
    debug: false,
    // if extraLng comes (ej. de Expo o RxDB), it has priority
    ...(extraLng ? { lng: extraLng } : {}),
  });
};

export default i18n;