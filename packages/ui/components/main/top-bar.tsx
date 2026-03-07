"use client";

import { Box } from "../ui/box";
import { HStack } from "../ui/hstack";
import { AtSignIcon, CheckIcon, CloseCircleIcon, InfoIcon } from "../ui/icon";
import { useAuth } from "../../auth/provider";
import { Alert, AlertIcon, AlertText } from "../ui/alert";
import { useRouter } from "solito/navigation";
import { useTranslation } from "react-i18next";
import { useAuthActions } from "../../auth/useAuthActions";
import { Button, ButtonIcon, ButtonText } from "../ui/button";
import { Tooltip, TooltipContent } from "../ui/tooltip";
import { Text } from "../ui/text";

const TopBar = ({ className }: { className?: string }) => {
    const { session, supabase } = useAuth()
    const { logout } = useAuthActions(supabase);
    const { replace } = useRouter()
    const { t } = useTranslation();

    const handleLogin = () => {
        replace('/auth/signup')
    };

    const handleLogout = () => {
        logout()
    };

    return (
        <HStack className={`p-2 ${className}`}>
            <Tooltip
                placement="bottom"
                trigger={(triggerProps) => {
                    return (
                        <Alert {...triggerProps} action={session ? "success" : "warning"} variant="solid" >
                            <AlertIcon size="xs" as={session ? CheckIcon : InfoIcon} />
                            <AlertText size="xs">{session
                                ? t('common.connected_as') + " " + session?.user?.email
                                : t('common.disconnected')}
                            </AlertText>
                        </Alert>
                    );
                }}>
                <TooltipContent className="p-4 rounded-md max-w-72 bg-background-50">
                    <Text className="font-sm text-justify">{t('common.offline_first_docs')}</Text>

                </TooltipContent>
            </Tooltip>
            <Box className="grow" />
            <Button size="xs" className="h-full" onPress={session ? handleLogout : handleLogin}>
                <ButtonIcon as={session ? CloseCircleIcon : AtSignIcon}></ButtonIcon>
                <ButtonText>{session
                    ? t('common.disconnect')
                    : t('common.connect')}
                </ButtonText>
            </Button>
        </HStack>
    );
};

export default TopBar;