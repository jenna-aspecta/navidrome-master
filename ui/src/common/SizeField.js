import React from 'react'
import PropTypes from 'prop-types'
import { formatBytes } from '../utils'
import { useRecordContext } from 'react-admin'
import { makeStyles } from '@material-ui/core'

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'inline-block',
  },
}))

export const SizeField = ({ source, ...rest }) => {
  const classes = useStyles()
  const record = useRecordContext(rest)
  return <span className={classes.root}>{formatBytes(record[source])}</span>
}

SizeField.propTypes = {
  label: PropTypes.string,
  record: PropTypes.object,
  source: PropTypes.string.isRequired,
}

SizeField.defaultProps = {
  addLabel: true,
}
