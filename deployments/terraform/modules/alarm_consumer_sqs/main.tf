resource "aws_cloudwatch_metric_alarm" "success-rate" {
  alarm_name          = "${var.family}-${var.application}-success-rate"
  count               = var.create ? 1 : 0
  datapoints_to_alarm = var.datapoints_to_alarm
  comparison_operator = "LessThanThreshold"
  evaluation_periods  = var.evaluation_periods
  threshold           = var.success_rate_threshold
  treat_missing_data  = "notBreaching"

  metric_query {
    id          = "messages"
    return_data = false

    metric {
      dimensions = {
        QueueName = "${var.project}-${var.environment}-${var.family}-${var.application}-${var.model}"
      }
      metric_name = "NumberOfMessagesReceived"
      namespace   = "AWS/SQS"
      period      = var.period
      stat        = "Sum"
    }
  }

  metric_query {
    id          = "errors"
    return_data = false

    metric {
      dimensions = {
        "reason" = "Error"
      }
      metric_name = "error"
      namespace   = "${var.project}/${var.environment}/${var.family}/${var.application}"
      period      = var.period
      stat        = "Sum"
    }
  }

  metric_query {
    expression  = "100-100*(errors/messages)"
    id          = "e1"
    label       = "100-100*(errors/messages)"
    return_data = true
  }
}
