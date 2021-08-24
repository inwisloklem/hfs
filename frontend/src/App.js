import * as React from 'react'

export function App() {
  const [hasError, setHasError] = React.useState(false)
  const [isLoadDisabled, setIsLoadDisabled] = React.useState(false)
  const [isSaveDisabled, setIsSaveDisabled] = React.useState(false)
  const [message, setMessage] = React.useState(null)

  const setHasErrorAndMessage = (message) => {
    setHasError(message.startsWith('Error'))
    setMessage(message)
  }

  const checkMessage = React.useCallback(
    () =>
      Promise.all([
        window.backend.Control.GetHasNoConfig(),
        window.backend.Control.GetMessage(),
        window.backend.Control.GetStoreDirSavesNumber(),
      ]).then(([hasNoConfig, message, savesNumber]) => {
        setHasErrorAndMessage(message)
        if (hasNoConfig) {
          setIsLoadDisabled(true)
          setIsSaveDisabled(true)
          return
        }
        setIsLoadDisabled(!savesNumber)
      }),
    []
  )

  React.useEffect(() => {
    checkMessage()
  }, [checkMessage])

  const handleLoadClick = () => {
    if (isSaveDisabled) {
      return
    }
    setIsLoadDisabled(true)
    window.backend.Control.LoadFile()
      .then(checkMessage)
      .then(() => setIsLoadDisabled(false))
  }

  const handleSaveClick = () => {
    if (isSaveDisabled) {
      return
    }
    setIsSaveDisabled(true)
    window.backend.Control.SaveFile()
      .then(checkMessage)
      .then(() => setIsSaveDisabled(false))
  }

  return (
    <div className="app">
      <div className={`message${hasError ? ' error' : ''}`}>{message}</div>
      <div className="container">
        <button
          className="primaryButton"
          disabled={isLoadDisabled}
          type="button"
          onClick={handleLoadClick}
        >
          Load last one
        </button>
        <button
          className="primaryButton"
          disabled={isSaveDisabled}
          type="button"
          onClick={handleSaveClick}
        >
          Save new entry
        </button>
      </div>
    </div>
  )
}
