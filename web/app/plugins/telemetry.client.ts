const ENDPOINT = 'https://nasak.ir/api/analytics/interactions'
const SAFE_TOKEN = /^[a-zA-Z0-9_.:-]{1,100}$/

function safeToken(value: string | null | undefined) {
  const normalized = value?.trim() || ''
  return SAFE_TOKEN.test(normalized) ? normalized : null
}

function sessionId() {
  try {
    const key = 'nasak-sms-analytics-session'
    const current = sessionStorage.getItem(key)
    if (current) return current
    const created = crypto.randomUUID()
    sessionStorage.setItem(key, created)
    return created
  } catch {
    return null
  }
}

export default defineNuxtPlugin(() => {
  const queue: Array<Record<string, unknown>> = []
  let timer: ReturnType<typeof setTimeout> | null = null
  const flush = () => {
    if (timer) clearTimeout(timer)
    timer = null
    if (!queue.length) return
    const payload = JSON.stringify({ events: queue.splice(0, 20) })
    if (
      navigator.sendBeacon?.(
        ENDPOINT,
        new Blob([payload], { type: 'application/json' }),
      )
    )
      return
    void fetch(ENDPOINT, {
      method: 'POST',
      headers: { 'content-type': 'application/json' },
      body: payload,
      keepalive: true,
    }).catch(() => undefined)
  }
  const onClick = (event: MouseEvent) => {
    if (!(event.target instanceof Element)) return
    const element =
      event.target.closest(
        'a,button,input,select,textarea,summary,[role],[data-analytics-id],[tabindex]',
      ) || event.target
    if (element.closest("[data-analytics-ignore='true']")) return
    const pointer = event as PointerEvent
    let href: string | null = null
    if (element instanceof HTMLAnchorElement) {
      try {
        const url = new URL(element.href, location.href)
        href = `${url.origin}${url.pathname}`.slice(0, 500)
      } catch {
        href = null
      }
    }
    queue.push({
      eventType: 'CLICK',
      path: location.pathname,
      sessionId: sessionId(),
      targetTag: element.tagName.toLowerCase(),
      targetId: safeToken(
        element.getAttribute('data-analytics-id') || element.id,
      ),
      targetName: safeToken(element.getAttribute('name')),
      targetRole: safeToken(element.getAttribute('role')),
      targetType:
        element instanceof HTMLInputElement ||
        element instanceof HTMLButtonElement
          ? safeToken(element.type)
          : null,
      href,
      occurredAt: new Date().toISOString(),
      metadata: {
        button: event.button,
        pointerType: safeToken(pointer.pointerType) || 'mouse',
        clientX: Math.round(event.clientX),
        clientY: Math.round(event.clientY),
        viewportWidth: innerWidth,
        viewportHeight: innerHeight,
      },
    })
    if (queue.length >= 20) flush()
    else if (!timer) timer = setTimeout(flush, 2500)
  }
  document.addEventListener('click', onClick, true)
  addEventListener('pagehide', flush)
})
