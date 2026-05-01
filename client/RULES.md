# Rules: Smart Parking Dashboard (LPR & IoT)

## Project Overview
A professional, high-performance monitoring dashboard for smart parking staff. Integrates License Plate Recognition (LPR) and IoT sensor data. Focus on real-time data visualization, clean aesthetics, and strict type safety.

## Tech Stack
- **Framework:** SvelteKit (Latest/Svelte 5 Runes)
- **Language:** TypeScript (Strict Mode)
- **Styling:** Tailwind CSS
- **UI Components:** shadcn-svelte
- **Icons:** Lucide Svelte
- **Charts:** LayerChart (Declarative Svelte SVG charts)

## 1. Core Architecture & Patterns
- **Atomic Components:** Breakdown UI into `atoms`, `molecules`, and `organisms` within `src/lib/components`.
- **Route-Based Logic:** Use `+page.svelte` for layout and `+page.ts/server.ts` for data fetching (smooth integration with the Golang REST API).
- **Runes Only:** Use `$state`, `$derived`, `$props`, and `$effect` for reactivity. Avoid the old `export let` syntax.
- **No Side Effects in Templates:** Keep logic inside `<script>` tags; templates should be declarative.
- **Build with a clean architecture**

## 2. Strict TypeScript Rules
- **Zero 'any' Policy:** Never use `any`. If a type is unknown, use `unknown` and type guards.
- **Interfaces & Types:** Define clear interfaces for LPR logs, Member data, and IoT sensor states mapping the Golang backend JSON responses.
- **Generic & Utility Types:** Maximize type reusability by utilizing Generic Types and TypeScript built-in utilities (`Pick`, `Omit`, `Partial`, `Record`, etc.) instead of declaring redundant new interfaces. Deriving types from existing base interfaces is highly preferred (e.g., using `Omit<Member, 'id'>` for a creation payload).
- **Zod Validation:** Use Zod for all form schemas (Login, Member CRUD) and API response validation.
- **Prop Typing:** All component props must be explicitly typed using TypeScript interfaces.

## 3. UI/UX, Styling & Animation Guidelines
- **Custom Color Palette:** You MUST configure Tailwind CSS theme variables (or shadcn base CSS) to use these exact hex codes:
  - **Primary:** `#2563EB` (Use for main buttons, primary actions, active navigation states).
  - **Secondary:** `#4F46E5` (Use for accents, highlights, secondary buttons, or chart elements).
  - **Tertiary:** `#BC4800` (Use for warnings, "OUT" status badges, or critical alerts).
  - **Neutral:** `#757681` (Use for secondary text, borders, subtle backgrounds, and inactive elements).
- **Clean Aesthetic:** Leverage the custom palette to maintain a professional, high-contrast interface. Avoid introducing external colors that clash with this scheme.
- **Consistent Spacing:** Follow a strict 4px/8px grid system using Tailwind classes.
- **Responsiveness (CRITICAL):** The UI MUST be fully responsive and optimized for all devices (mobile, tablet, desktop). Use Tailwind's responsive breakpoints (`sm:`, `md:`, `lg:`, `xl:`) appropriately. Ensure sidebars collapse into hamburger menus on mobile, and complex tables have horizontal scrolling or stack gracefully.
- **Animations (Svelte Native):** DO NOT use Framer Motion. Use Svelte's built-in `svelte/transition` and `svelte/motion`. Animations must be subtle and functional:
  - Use `spring` or `tweened` for Live Counters (e.g., smoothly animating the "Currently Inside" number).
  - Use `fly` or `slide` for dynamically inserting new rows into the "Recent Scans" list.
  - Use shadcn `Skeleton` (pulse effect) for loading states while fetching data from the Golang backend.
- **Accessibility (a11y):** Ensure proper ARIA labels, keyboard navigation for tables, and high contrast ratios.

## 4. Page Specific Requirements

### Login Page
- Clean, centered card.
- Client-side validation before submission.
- Focus on security and simplicity. Use Primary color for the main call-to-action.
- Must be perfectly centered and padded on small mobile screens.

### Dashboard (Monitoring)
- **Stats Grid:** 4 columns for Live Count, Entry, Exit, and Availability. (Must collapse to 1 or 2 columns on mobile/tablet).
- **Visuals:** Use LayerChart to create smooth Line/Area charts for traffic flow utilizing Primary and Secondary colors.
- **Real-time Vibe:** The "Recent Scans" list should feel dynamic with smooth entry animations.

### Member Management
- Implement a robust `DataTable` with sorting and searching.
- Use shadcn `Dialog` for adding/editing members to keep the user in context.

### Activity Logs (Gate History)
- **Complex Filtering:** Implement logic for Status (IN/OUT) and Date Ranges. Make sure the filter bar wraps neatly on mobile.
- **Status Badges:** Explicitly use `#2563EB` (Primary) or a specific green for "IN" and `#BC4800` (Tertiary) for "OUT" to make scanning logs visually instantaneous.
- **Search:** Instant search across license plates.
- **Performance:** Optimized rendering for large log lists (pagination or virtualization).

## 5. Coding Standards
- **Naming:** `camelCase` for variables/functions, `PascalCase` for components, `kebab-case` for CSS classes.
- **File Structure:** Keep `src/lib` organized:
  - `/components/ui`: Shadcn primitives.
  - `/components/shared`: Custom reusable components.
  - `/types`: All TS interfaces.
  - `/utils`: Helper functions and formatters.
- **DRY:** Extract repetitive logic (like date formatting or status color mapping) into utility functions.