import Select from 'react-select';

import i18n, { changeLocaleLanguage } from '@/modules/i18n/config';
import { useToast } from '@/domain/common/hooks/useToast';

export const LocaleLanguageSelect = () => {
  const toast = useToast();
  const options = [
    { label: 'Japanese', value: 'ja' },
    { label: 'English', value: 'en' },
  ] as const;
  const defaultValue = options.find(opt => opt.value === i18n.language);
  return (
    <Select
      options={options}
      defaultValue={defaultValue}
      onChange={v => {
        if (!v) {
          return;
        }
        changeLocaleLanguage(v.value);
        toast.changedLocaleLanguage();
      }}
    />
  );
};
