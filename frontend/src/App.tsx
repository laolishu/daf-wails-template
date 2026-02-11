import {
  AppShell,
  ActionIcon,
  Box,
  Divider,
  Group,
  Image,
  NavLink,
  Stack,
  Text,
} from "@mantine/core";
import {
  IconHeart,
  IconArrowUpRight,
  IconHelpCircle,
  IconSettings,
  IconLayoutGrid,
  IconMinus,
  IconSquare,
  IconX,
} from "@tabler/icons-react";
import { useEffect, useState } from "react";
import { Quit, WindowMinimise, WindowToggleMaximise } from "../wailsjs/runtime";
import HealthCheck from "./pages/HealthCheck";
import ExamplesPage from "./pages/Examples";
import SettingsPage from "./pages/Settings";
import HelpPage from "./pages/Help";
import UpgradePage from "./pages/Upgrade";
import { fetchWindowTitle } from "./services/windowTitle";
import { fetchLanguage, updateLanguage } from "./services/language";
import { messages, normalizeLanguage, type Language } from "./i18n";

export default function App() {
  const colors = {
    mainBg: "#F6F8FB",
    sidebarBg: "#3A475A",
    sidebarText: "#C9D1DC",
    sidebarTextActive: "#FFFFFF",
    sidebarHoverBg: "rgba(255,255,255,0.08)",
    sidebarActiveBg: "rgba(255,255,255,0.12)",
    accentBlue: "#3B82F6",
  };
  const [activePage, setActivePage] = useState<
    "health" | "examples" | "settings" | "help" | "upgrade"
  >("health");
  const [appTitle, setAppTitle] = useState("DAF Wails App");
  const [language, setLanguage] = useState<Language>("zh-CN");
  const primaryNavItems = [
    { label: messages[language].navHealthCheck, page: "health" as const, icon: IconHeart },
    { label: messages[language].navExamples, page: "examples" as const, icon: IconLayoutGrid },
  ];
  const secondaryNavItems = [
    { label: messages[language].navSettings, page: "settings" as const, icon: IconSettings },
    { label: messages[language].navHelp, page: "help" as const, icon: IconHelpCircle },
    { label: messages[language].navUpgrade, page: "upgrade" as const, icon: IconArrowUpRight },
  ];

  useEffect(() => {
    fetchWindowTitle()
      .then((title) => setAppTitle(title || "DAF Wails App"))
      .catch(() => setAppTitle("DAF Wails App"));
  }, []);

  useEffect(() => {
    fetchLanguage()
      .then((value) => setLanguage(normalizeLanguage(value)))
      .catch(() => setLanguage("zh-CN"));
  }, []);

  const handleLanguageChange = async (value: Language) => {
    const previous = language;
    setLanguage(value);
    try {
      await updateLanguage(value);
    } catch {
      setLanguage(previous);
    }
  };

  const handleTitlebarDoubleClick = (
    event: React.MouseEvent<HTMLElement>
  ) => {
    const target = event.target as HTMLElement | null;
    if (target?.closest(".wails-no-drag")) {
      return;
    }
    WindowToggleMaximise();
  };

  return (
    <AppShell
      padding={0}
      layout="default"
      navbar={{ width: 208, breakpoint: 0, collapsed: { mobile: false, desktop: false } }}
      header={{ height: 40 }}
      styles={{
        root: {
          height: "100%",
        },
        navbar: {
          backgroundColor: colors.sidebarBg,
          color: colors.sidebarText,
          borderRight: 0,
        },
        main: {
          backgroundColor: colors.mainBg,
          height: "100%",
          overflow: "hidden",
        },
        header: {
          backgroundColor: colors.sidebarBg,
          borderBottom: 0,
        },
      }}
    >
      <AppShell.Header
        className="wails-drag"
        onDoubleClick={handleTitlebarDoubleClick}
      >
        <Group h="100%" px="lg" justify="space-between" w="100%">
          <Group gap="sm">
            <Image src="/logo.png" w={20} h={20} fit="contain" />
            <Text
              fw={600}
              c={colors.sidebarTextActive}
              lineClamp={1}
              style={{ whiteSpace: "nowrap" }}
            >
              {appTitle}
            </Text>
          </Group>
          <Group gap={4} className="wails-no-drag">
            <ActionIcon
              variant="subtle"
              color="white"
              radius="sm"
              size={28}
              className="wails-no-drag"
              onClick={() => WindowMinimise()}
            >
              <IconMinus size={16} color="#ffffff" />
            </ActionIcon>
            <ActionIcon
              variant="subtle"
              color="white"
              radius="sm"
              size={28}
              className="wails-no-drag"
              onClick={() => WindowToggleMaximise()}
            >
              <IconSquare size={14} color="#ffffff" />
            </ActionIcon>
            <ActionIcon
              variant="subtle"
              color="white"
              radius="sm"
              size={28}
              className="wails-no-drag"
              onClick={() => Quit()}
            >
              <IconX size={16} color="#ffffff" />
            </ActionIcon>
          </Group>
        </Group>
      </AppShell.Header>
      <AppShell.Navbar>
        <Stack h="100%" p="md" gap="xl" justify="space-between">
          <Stack gap="lg">
            <Stack gap={8}>
              {primaryNavItems.map((item) => {
                const Icon = item.icon;
                const isActive = item.page && activePage === item.page;
                const iconColor = isActive
                  ? colors.sidebarTextActive
                  : colors.sidebarText;
                return (
                  <NavLink
                    key={item.label}
                    label={item.label}
                    leftSection={<Icon size={18} color={iconColor} />}
                    active={isActive}
                    onClick={() => item.page && setActivePage(item.page)}
                    styles={{
                      root: {
                        borderRadius: 8,
                        height: 40,
                        padding: "6px 10px",
                        display: "flex",
                        alignItems: "center",
                        columnGap: 10,
                        backgroundColor: isActive
                          ? colors.sidebarActiveBg
                          : "transparent",
                        color: isActive
                          ? colors.sidebarTextActive
                          : colors.sidebarText,
                        cursor: "pointer",
                        boxShadow: "none",
                        "&:hover": {
                          backgroundColor: colors.sidebarHoverBg,
                          color: colors.sidebarTextActive,
                        },
                      },
                      section: {
                        marginInlineStart: 6,
                        marginInlineEnd: 10,
                      },
                      label: {
                        fontSize: 14,
                        fontWeight: 600,
                        whiteSpace: "nowrap",
                        overflow: "hidden",
                        textOverflow: "ellipsis",
                      },
                    }}
                  />
                );
              })}
            </Stack>
          </Stack>
          <Stack gap="md">
            <Divider color="rgba(255,255,255,0.16)" />
            <Stack gap={8}>
              {secondaryNavItems.map((item) => {
                const Icon = item.icon;
                const isActive = item.page && activePage === item.page;
                const iconColor = isActive
                  ? colors.sidebarTextActive
                  : colors.sidebarText;
                return (
                  <NavLink
                    key={item.label}
                    label={item.label}
                    leftSection={<Icon size={18} color={iconColor} />}
                    active={isActive}
                    onClick={() => item.page && setActivePage(item.page)}
                    styles={{
                      root: {
                        borderRadius: 8,
                        height: 40,
                        padding: "6px 10px",
                        display: "flex",
                        alignItems: "center",
                        columnGap: 10,
                        backgroundColor: isActive
                          ? colors.sidebarActiveBg
                          : "transparent",
                        color: isActive
                          ? colors.sidebarTextActive
                          : colors.sidebarText,
                        cursor: "pointer",
                        boxShadow: "none",
                        "&:hover": {
                          backgroundColor: colors.sidebarHoverBg,
                          color: colors.sidebarTextActive,
                        },
                      },
                      section: {
                        marginInlineStart: 6,
                        marginInlineEnd: 10,
                      },
                      label: {
                        fontSize: 14,
                        fontWeight: 600,
                        whiteSpace: "nowrap",
                        overflow: "hidden",
                        textOverflow: "ellipsis",
                      },
                    }}
                  />
                );
              })}
            </Stack>
          </Stack>
        </Stack>
      </AppShell.Navbar>
      <AppShell.Main>
        <Box
          w="100%"
          h="100%"
          p="xl"
          style={{
            display: "flex",
            flexDirection: "column",
            alignItems: "stretch",
            overflow: "auto",
            borderRadius: 12,
          }}
        >
          <Box w="100%" style={{ flex: 1 }}>
            {activePage === "health" && (
              <HealthCheck messages={messages[language]} />
            )}
            {activePage === "examples" && (
              <ExamplesPage messages={messages[language]} />
            )}
            {activePage === "settings" && (
              <SettingsPage
                language={language}
                messages={messages[language]}
                onLanguageChange={handleLanguageChange}
              />
            )}
            {activePage === "help" && <HelpPage messages={messages[language]} />}
            {activePage === "upgrade" && (
              <UpgradePage messages={messages[language]} />
            )}
          </Box>
        </Box>
      </AppShell.Main>
    </AppShell>
  );
}
