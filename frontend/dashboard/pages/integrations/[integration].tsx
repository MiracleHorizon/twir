import { LoadingOverlay } from '@mantine/core';
import { serverSideTranslations } from 'next-i18next/serverSideTranslations';
import { useRouter } from 'next/router';
import { useContext, useEffect } from 'react';

import { useDonationAlerts, useFaceit, useLastfm, useSpotify, useStreamlabs, useVK, type Integration as IntegrationType } from '@/services/api/integrations';
import { SelectedDashboardContext } from '@/services/selectedDashboardProvider';

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
export const getServerSideProps = async ({ locale }) => ({
  props: {
    ...(await serverSideTranslations(locale, ['integrations', 'layout'])),
  },
});

export default function Integration() {
  const router = useRouter();

  const managers: Record<string, IntegrationType<any>> = {
    'donationalerts': useDonationAlerts(),
    'faceit': useFaceit(),
    'lastfm': useLastfm(),
    'spotify': useSpotify(),
    'streamlabs': useStreamlabs(),
    'vk': useVK(),
  };
  const { integration } = router.query;
  const { usePostCode } = managers[integration as string];
  const poster = usePostCode();

  useEffect(() => {
    const { code, token } = router.query;

    const incomingCode = code ?? token;

    if (typeof incomingCode !== 'string' || typeof integration !== 'string' || !(integration in managers)) {
      router.push('/integrations');
      return;
    }

    poster.mutateAsync({ code: incomingCode }).finally(() => {
      router.push('/integrations');
    });
  }, []);

  return <LoadingOverlay visible={true} overlayBlur={2} />;
}