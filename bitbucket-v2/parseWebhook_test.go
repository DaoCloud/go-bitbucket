package bitbucket_v2

import (
	"encoding/json"
	"testing"
)

const (
	pullrequest = `{
  "pullrequest": {
    "merge_commit": null,
    "description": "test",
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/test2/pullrequests/1"
      },
      "html": {
        "href": "https://bitbucket.org/xuzhenglun/test2/pull-requests/1"
      }
    },
    "title": "Master",
    "close_source_branch": false,
    "reviewers": [],
    "destination": {
      "commit": {
        "hash": "58e2d50c6fba",
        "links": {
          "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/test2/commit/58e2d50c6fba"
          }
        }
      },
      "repository": {
        "links": {
          "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/test2"
          },
          "html": {
            "href": "https://bitbucket.org/xuzhenglun/test2"
          },
          "avatar": {
            "href": "https://bitbucket.org/xuzhenglun/test2/avatar/32/"
          }
        },
        "type": "repository",
        "name": "test2",
        "full_name": "xuzhenglun/test2",
        "uuid": "{550f38e2-d2e3-4d2a-9a58-79c359cafbed}"
      },
      "branch": {
        "name": "branch2"
      }
    },
    "comment_count": 0,
    "closed_by": null,
    "source": {
      "commit": {
        "hash": "b6d250af77e4",
        "links": {
          "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/test2/commit/b6d250af77e4"
          }
        }
      },
      "repository": {
        "links": {
          "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/test2"
          },
          "html": {
            "href": "https://bitbucket.org/xuzhenglun/test2"
          },
          "avatar": {
            "href": "https://bitbucket.org/xuzhenglun/test2/avatar/32/"
          }
        },
        "type": "repository",
        "name": "test2",
        "full_name": "xuzhenglun/test2",
        "uuid": "{550f38e2-d2e3-4d2a-9a58-79c359cafbed}"
      },
      "branch": {
        "name": "master"
      }
    },
    "created_on": "2016-06-14T05:51:27.848161+00:00",
    "state": "OPEN",
    "task_count": 0,
    "participants": [],
    "reason": "",
    "updated_on": "2016-06-14T05:51:27.876309+00:00",
    "author": {
      "username": "xuzhenglun",
      "display_name": "xuzhenglun",
      "type": "user",
      "uuid": "{6274aeff-436b-49d0-85c4-8ac868aa6dbb}",
      "links": {
        "self": {
          "href": "https://api.bitbucket.org/2.0/users/xuzhenglun"
        },
        "html": {
          "href": "https://bitbucket.org/xuzhenglun/"
        },
        "avatar": {
          "href": "https://bitbucket.org/account/xuzhenglun/avatar/32/"
        }
      }
    },
    "type": "pullrequest",
    "id": 1
  },
  "actor": {
    "username": "xuzhenglun",
    "display_name": "xuzhenglun",
    "type": "user",
    "uuid": "{6274aeff-436b-49d0-85c4-8ac868aa6dbb}",
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/users/xuzhenglun"
      },
      "html": {
        "href": "https://bitbucket.org/xuzhenglun/"
      },
      "avatar": {
        "href": "https://bitbucket.org/account/xuzhenglun/avatar/32/"
      }
    }
  },
  "repository": {
    "scm": "git",
    "website": "",
    "name": "test2",
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/test2"
      },
      "html": {
        "href": "https://bitbucket.org/xuzhenglun/test2"
      },
      "avatar": {
        "href": "https://bitbucket.org/xuzhenglun/test2/avatar/32/"
      }
    },
    "full_name": "xuzhenglun/test2",
    "owner": {
      "username": "xuzhenglun",
      "display_name": "xuzhenglun",
      "type": "user",
      "uuid": "{6274aeff-436b-49d0-85c4-8ac868aa6dbb}",
      "links": {
        "self": {
          "href": "https://api.bitbucket.org/2.0/users/xuzhenglun"
        },
        "html": {
          "href": "https://bitbucket.org/xuzhenglun/"
        },
        "avatar": {
          "href": "https://bitbucket.org/account/xuzhenglun/avatar/32/"
        }
      }
    },
    "type": "repository",
    "is_private": false,
    "uuid": "{550f38e2-d2e3-4d2a-9a58-79c359cafbed}"
  }
}`

	push = `{
  "push": {
    "changes": [
      {
        "forced": false,
        "old": {
          "links": {
            "commits": {
              "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/docker-testing-docker-deployment/commits/master"
            },
            "self": {
              "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/docker-testing-docker-deployment/refs/branches/master"
            },
            "html": {
              "href": "https://bitbucket.org/xuzhenglun/docker-testing-docker-deployment/branch/master"
            }
          },
          "type": "branch",
          "target": {
            "hash": "e4aac047fcfed2d9e0aeb200e93617c388b1cad9",
            "links": {
              "self": {
                "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/docker-testing-docker-deployment/commit/e4aac047fcfed2d9e0aeb200e93617c388b1cad9"
              },
              "html": {
                "href": "https://bitbucket.org/xuzhenglun/docker-testing-docker-deployment/commits/e4aac047fcfed2d9e0aeb200e93617c388b1cad9"
              }
            },
            "author": {
              "raw": "Reficul <xuzhenglun@gmail.com>",
              "user": {
                "username": "xuzhenglun",
                "type": "user",
                "display_name": "xuzhenglun",
                "uuid": "{6274aeff-436b-49d0-85c4-8ac868aa6dbb}",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/users/xuzhenglun"
                  },
                  "html": {
                    "href": "https://bitbucket.org/xuzhenglun/"
                  },
                  "avatar": {
                    "href": "https://bitbucket.org/account/xuzhenglun/avatar/32/"
                  }
                }
              }
            },
            "parents": [
              {
                "type": "commit",
                "hash": "12cbcedfd8502a4bb12d3b955200cf3c833ba8a3",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/docker-testing-docker-deployment/commit/12cbcedfd8502a4bb12d3b955200cf3c833ba8a3"
                  },
                  "html": {
                    "href": "https://bitbucket.org/xuzhenglun/docker-testing-docker-deployment/commits/12cbcedfd8502a4bb12d3b955200cf3c833ba8a3"
                  }
                }
              }
            ],
            "date": "2016-06-14T03:31:19+00:00",
            "message": "this is a hook test\n",
            "type": "commit"
          },
          "repository": {
            "links": {
              "self": {
                "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/docker-testing-docker-deployment"
              },
              "html": {
                "href": "https://bitbucket.org/xuzhenglun/docker-testing-docker-deployment"
              },
              "avatar": {
                "href": "https://bitbucket.org/xuzhenglun/docker-testing-docker-deployment/avatar/32/"
              }
            },
            "type": "repository",
            "uuid": "{4364edb7-ac79-4fd9-896d-59d994cd89a1}",
            "full_name": "xuzhenglun/docker-testing-docker-deployment",
            "name": "docker-testing-docker-deployment"
          },
          "name": "master"
        },
        "links": {
          "commits": {
            "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/docker-testing-docker-deployment/commits?include=8a4e3a7f4a6510d9edbd2c1caef17bc238224f87&exclude=e4aac047fcfed2d9e0aeb200e93617c388b1cad9"
          },
          "html": {
            "href": "https://bitbucket.org/xuzhenglun/docker-testing-docker-deployment/branches/compare/8a4e3a7f4a6510d9edbd2c1caef17bc238224f87..e4aac047fcfed2d9e0aeb200e93617c388b1cad9"
          },
          "diff": {
            "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/docker-testing-docker-deployment/diff/8a4e3a7f4a6510d9edbd2c1caef17bc238224f87..e4aac047fcfed2d9e0aeb200e93617c388b1cad9"
          }
        },
        "truncated": false,
        "commits": [
          {
            "hash": "8a4e3a7f4a6510d9edbd2c1caef17bc238224f87",
            "links": {
              "self": {
                "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/docker-testing-docker-deployment/commit/8a4e3a7f4a6510d9edbd2c1caef17bc238224f87"
              },
              "html": {
                "href": "https://bitbucket.org/xuzhenglun/docker-testing-docker-deployment/commits/8a4e3a7f4a6510d9edbd2c1caef17bc238224f87"
              }
            },
            "author": {
              "raw": "Reficul <xuzhenglun@gmail.com>",
              "user": {
                "username": "xuzhenglun",
                "type": "user",
                "display_name": "xuzhenglun",
                "uuid": "{6274aeff-436b-49d0-85c4-8ac868aa6dbb}",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/users/xuzhenglun"
                  },
                  "html": {
                    "href": "https://bitbucket.org/xuzhenglun/"
                  },
                  "avatar": {
                    "href": "https://bitbucket.org/account/xuzhenglun/avatar/32/"
                  }
                }
              }
            },
            "parents": [
              {
                "type": "commit",
                "hash": "e4aac047fcfed2d9e0aeb200e93617c388b1cad9",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/docker-testing-docker-deployment/commit/e4aac047fcfed2d9e0aeb200e93617c388b1cad9"
                  },
                  "html": {
                    "href": "https://bitbucket.org/xuzhenglun/docker-testing-docker-deployment/commits/e4aac047fcfed2d9e0aeb200e93617c388b1cad9"
                  }
                }
              }
            ],
            "date": "2016-06-14T06:11:45+00:00",
            "message": "this is a hook test 2\n",
            "type": "commit"
          }
        ],
        "created": false,
        "closed": false,
        "new": {
          "links": {
            "commits": {
              "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/docker-testing-docker-deployment/commits/master"
            },
            "self": {
              "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/docker-testing-docker-deployment/refs/branches/master"
            },
            "html": {
              "href": "https://bitbucket.org/xuzhenglun/docker-testing-docker-deployment/branch/master"
            }
          },
          "type": "branch",
          "target": {
            "hash": "8a4e3a7f4a6510d9edbd2c1caef17bc238224f87",
            "links": {
              "self": {
                "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/docker-testing-docker-deployment/commit/8a4e3a7f4a6510d9edbd2c1caef17bc238224f87"
              },
              "html": {
                "href": "https://bitbucket.org/xuzhenglun/docker-testing-docker-deployment/commits/8a4e3a7f4a6510d9edbd2c1caef17bc238224f87"
              }
            },
            "author": {
              "raw": "Reficul <xuzhenglun@gmail.com>",
              "user": {
                "username": "xuzhenglun",
                "type": "user",
                "display_name": "xuzhenglun",
                "uuid": "{6274aeff-436b-49d0-85c4-8ac868aa6dbb}",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/users/xuzhenglun"
                  },
                  "html": {
                    "href": "https://bitbucket.org/xuzhenglun/"
                  },
                  "avatar": {
                    "href": "https://bitbucket.org/account/xuzhenglun/avatar/32/"
                  }
                }
              }
            },
            "parents": [
              {
                "type": "commit",
                "hash": "e4aac047fcfed2d9e0aeb200e93617c388b1cad9",
                "links": {
                  "self": {
                    "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/docker-testing-docker-deployment/commit/e4aac047fcfed2d9e0aeb200e93617c388b1cad9"
                  },
                  "html": {
                    "href": "https://bitbucket.org/xuzhenglun/docker-testing-docker-deployment/commits/e4aac047fcfed2d9e0aeb200e93617c388b1cad9"
                  }
                }
              }
            ],
            "date": "2016-06-14T06:11:45+00:00",
            "message": "this is a hook test 2\n",
            "type": "commit"
          },
          "repository": {
            "links": {
              "self": {
                "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/docker-testing-docker-deployment"
              },
              "html": {
                "href": "https://bitbucket.org/xuzhenglun/docker-testing-docker-deployment"
              },
              "avatar": {
                "href": "https://bitbucket.org/xuzhenglun/docker-testing-docker-deployment/avatar/32/"
              }
            },
            "type": "repository",
            "uuid": "{4364edb7-ac79-4fd9-896d-59d994cd89a1}",
            "full_name": "xuzhenglun/docker-testing-docker-deployment",
            "name": "docker-testing-docker-deployment"
          },
          "name": "master"
        }
      }
    ]
  },
  "actor": {
    "username": "xuzhenglun",
    "type": "user",
    "display_name": "xuzhenglun",
    "uuid": "{6274aeff-436b-49d0-85c4-8ac868aa6dbb}",
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/users/xuzhenglun"
      },
      "html": {
        "href": "https://bitbucket.org/xuzhenglun/"
      },
      "avatar": {
        "href": "https://bitbucket.org/account/xuzhenglun/avatar/32/"
      }
    }
  },
  "repository": {
    "scm": "git",
    "website": "",
    "name": "docker-testing-docker-deployment",
    "links": {
      "self": {
        "href": "https://api.bitbucket.org/2.0/repositories/xuzhenglun/docker-testing-docker-deployment"
      },
      "html": {
        "href": "https://bitbucket.org/xuzhenglun/docker-testing-docker-deployment"
      },
      "avatar": {
        "href": "https://bitbucket.org/xuzhenglun/docker-testing-docker-deployment/avatar/32/"
      }
    },
    "full_name": "xuzhenglun/docker-testing-docker-deployment",
    "owner": {
      "username": "xuzhenglun",
      "type": "user",
      "display_name": "xuzhenglun",
      "uuid": "{6274aeff-436b-49d0-85c4-8ac868aa6dbb}",
      "links": {
        "self": {
          "href": "https://api.bitbucket.org/2.0/users/xuzhenglun"
        },
        "html": {
          "href": "https://bitbucket.org/xuzhenglun/"
        },
        "avatar": {
          "href": "https://bitbucket.org/account/xuzhenglun/avatar/32/"
        }
      }
    },
    "type": "repository",
    "is_private": true,
    "uuid": "{4364edb7-ac79-4fd9-896d-59d994cd89a1}"
  }
}`
)

func Test_ParsePRHook(T *testing.T) {
	hook, err := ParsePRHook([]byte(pullrequest))
	if err != nil {
		T.Error(err)
		return
	}
	j, _ := json.Marshal(hook)
	T.Log(string(j))
}

func Test_ParsePushHook(T *testing.T) {
	hook, err := ParseRepoHook([]byte(push))
	if err != nil {
		T.Error(err)
		return
	}
	j, _ := json.Marshal(hook)
	T.Log(string(j))
}
