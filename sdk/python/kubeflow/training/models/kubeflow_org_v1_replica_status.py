# coding: utf-8

"""
    Kubeflow Training SDK

    Python SDK for Kubeflow Training  # noqa: E501

    The version of the OpenAPI document: v1.9.0
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from kubeflow.training.configuration import Configuration


class KubeflowOrgV1ReplicaStatus(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'active': 'int',
        'failed': 'int',
        'label_selector': 'V1LabelSelector',
        'selector': 'str',
        'succeeded': 'int'
    }

    attribute_map = {
        'active': 'active',
        'failed': 'failed',
        'label_selector': 'labelSelector',
        'selector': 'selector',
        'succeeded': 'succeeded'
    }

    def __init__(self, active=None, failed=None, label_selector=None, selector=None, succeeded=None, local_vars_configuration=None):  # noqa: E501
        """KubeflowOrgV1ReplicaStatus - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._active = None
        self._failed = None
        self._label_selector = None
        self._selector = None
        self._succeeded = None
        self.discriminator = None

        if active is not None:
            self.active = active
        if failed is not None:
            self.failed = failed
        if label_selector is not None:
            self.label_selector = label_selector
        if selector is not None:
            self.selector = selector
        if succeeded is not None:
            self.succeeded = succeeded

    @property
    def active(self):
        """Gets the active of this KubeflowOrgV1ReplicaStatus.  # noqa: E501

        The number of actively running pods.  # noqa: E501

        :return: The active of this KubeflowOrgV1ReplicaStatus.  # noqa: E501
        :rtype: int
        """
        return self._active

    @active.setter
    def active(self, active):
        """Sets the active of this KubeflowOrgV1ReplicaStatus.

        The number of actively running pods.  # noqa: E501

        :param active: The active of this KubeflowOrgV1ReplicaStatus.  # noqa: E501
        :type: int
        """

        self._active = active

    @property
    def failed(self):
        """Gets the failed of this KubeflowOrgV1ReplicaStatus.  # noqa: E501

        The number of pods which reached phase Failed.  # noqa: E501

        :return: The failed of this KubeflowOrgV1ReplicaStatus.  # noqa: E501
        :rtype: int
        """
        return self._failed

    @failed.setter
    def failed(self, failed):
        """Sets the failed of this KubeflowOrgV1ReplicaStatus.

        The number of pods which reached phase Failed.  # noqa: E501

        :param failed: The failed of this KubeflowOrgV1ReplicaStatus.  # noqa: E501
        :type: int
        """

        self._failed = failed

    @property
    def label_selector(self):
        """Gets the label_selector of this KubeflowOrgV1ReplicaStatus.  # noqa: E501


        :return: The label_selector of this KubeflowOrgV1ReplicaStatus.  # noqa: E501
        :rtype: V1LabelSelector
        """
        return self._label_selector

    @label_selector.setter
    def label_selector(self, label_selector):
        """Sets the label_selector of this KubeflowOrgV1ReplicaStatus.


        :param label_selector: The label_selector of this KubeflowOrgV1ReplicaStatus.  # noqa: E501
        :type: V1LabelSelector
        """

        self._label_selector = label_selector

    @property
    def selector(self):
        """Gets the selector of this KubeflowOrgV1ReplicaStatus.  # noqa: E501

        A Selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty Selector matches all objects. A null Selector matches no objects.  # noqa: E501

        :return: The selector of this KubeflowOrgV1ReplicaStatus.  # noqa: E501
        :rtype: str
        """
        return self._selector

    @selector.setter
    def selector(self, selector):
        """Sets the selector of this KubeflowOrgV1ReplicaStatus.

        A Selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty Selector matches all objects. A null Selector matches no objects.  # noqa: E501

        :param selector: The selector of this KubeflowOrgV1ReplicaStatus.  # noqa: E501
        :type: str
        """

        self._selector = selector

    @property
    def succeeded(self):
        """Gets the succeeded of this KubeflowOrgV1ReplicaStatus.  # noqa: E501

        The number of pods which reached phase Succeeded.  # noqa: E501

        :return: The succeeded of this KubeflowOrgV1ReplicaStatus.  # noqa: E501
        :rtype: int
        """
        return self._succeeded

    @succeeded.setter
    def succeeded(self, succeeded):
        """Sets the succeeded of this KubeflowOrgV1ReplicaStatus.

        The number of pods which reached phase Succeeded.  # noqa: E501

        :param succeeded: The succeeded of this KubeflowOrgV1ReplicaStatus.  # noqa: E501
        :type: int
        """

        self._succeeded = succeeded

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, KubeflowOrgV1ReplicaStatus):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, KubeflowOrgV1ReplicaStatus):
            return True

        return self.to_dict() != other.to_dict()
