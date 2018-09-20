import axios, { AxiosResponse, AxiosError } from 'axios'
import { resolve } from 'url'
import { rejects } from 'assert'
import { Member, User } from './types'
import { log } from 'util'

const instance = axios.create({
  baseURL: `${process.env.VUE_APP_BASEURL}/api/`,
})

const get = (url: string): Promise<any> => {
  return new Promise((rsv, rjt) => {
    instance
      .get(url)
      .then((res: AxiosResponse) => {
        rsv(res.data)
      })
      .catch((err: AxiosError) => {
        rjt(null)
      })
  })
}

const post = (url: string, u: User): Promise<any> => {
  return new Promise((rsv, rjt) => {
    instance
      .post('user', u)
      .then((res: AxiosResponse) => {
        rsv()
      })
      .catch((err: AxiosError) => {
        rjt()
      })
  })
}

export const ais = { get, post }

export const waitSecond = (sec: number): Promise<any> => {
  return new Promise((rsv) => {
    const id = setTimeout(() => {
      rsv(id)
    }, sec * 1000)
  })
}

export const initNames = (ms: Member[]): string[] => {
  const names: string[] = []
  const cms: Member[] = Object.assign([], ms)
  while (cms.length !== 0) {
    const r = Math.floor(Math.random() * cms.length)
    names.push(cms[r].name)
    cms.splice(r, 1)
  }
  return names
}

export const initMembers = (ms: Member[]): Member[] => {
  for (let i = ms.length - 1; i > 0; i--) {
    const r = Math.floor(Math.random() * (i + 1))
    const tmp = ms[i]
    ms[i] = ms[r]
    ms[r] = tmp
  }
  return ms
}
