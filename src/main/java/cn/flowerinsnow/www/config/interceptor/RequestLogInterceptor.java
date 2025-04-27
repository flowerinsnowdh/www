/*
 * www
 * Copyright (C) 2025  flowerinsnow
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
package cn.flowerinsnow.www.config.interceptor;

import cn.flowerinsnow.www.mapper.AccessMapper;
import com.google.gson.Gson;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;

@Component
public class RequestLogInterceptor implements HandlerInterceptor {
    @Autowired
    private AccessMapper accessMapper;

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) {
        String remoteAddress = request.getHeader("X-Real-IP");
        if (!remoteAddress.isEmpty()) {
            this.accessMapper.insert(
                    request.getMethod(),
                    request.getHeader("Host"),
                    request.getServletPath(),
                    new Gson().toJson(request.getParameterMap()),
                    request.getHeader("User-Agent"),
                    request.getHeader("referer"),
                    remoteAddress
            );
        }
        return true;
    }
}
